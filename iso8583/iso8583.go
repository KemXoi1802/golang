package iso8583

import (
	"encoding/hex"
	"errors"
	"fmt"
	"golang/logging"
	"sort"
	"strconv"
)

var bitCheck = [8]byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}

//bitAttribute
type bitAttribute struct {
	mBitType         BitType
	mLengthAttribute BitLength
	mMaxLen          int
	mLen             int
	mData            []byte
	mIsSet           bool
}

//Message store data of iso
type Message struct {
	mTPDU         []byte
	mMTI          int
	mBitMap       []byte
	mFieldEnabled []int
	mFieldsAttrs  map[int]bitAttribute
	mPackageSize  int
	mBuffer       []byte
}

const (
	maxBufferSize = 4096
	maxTpduSize   = 5
	maxBitmapSize = 8
)

//DefaultIso8583Data init
func DefaultIso8583Data() *Message {
	return &Message{
		mTPDU:        make([]byte, maxTpduSize),
		mBitMap:      make([]byte, maxBitmapSize),
		mFieldsAttrs: make(map[int]bitAttribute),
		mBuffer:      make([]byte, maxBufferSize),
	}
}

//NewIso8583Data init
func NewIso8583Data(data []byte, size int) *Message {
	var mti int
	var err error
	if mti, err = strconv.Atoi(hex.EncodeToString(data[5:7])); err != nil {
		logging.GetLog().Info("MTI is not type int")
	}
	// if len(data) != size {
	// 	logging.GetLog().Info("DATA or LENGTH invalid")
	// }
	return &Message{
		//tpdu, mti, bitmap.
		mTPDU:   data[:5],
		mMTI:    mti,
		mBitMap: data[7:15],
		//start from MTI
		mBuffer:      data[5:size],
		mFieldsAttrs: make(map[int]bitAttribute),
		mPackageSize: size,
	}
}

//PackField pack data into field
func (i *Message) PackField(FieldID int, FieldData string) {
	FieldAttr := Spec[FieldID]

	switch FieldAttr.FieldType {
	case An, Ans:
		if data, ok := StringToAsc(FieldData); ok == nil {
			i.mFieldsAttrs[FieldID] = bitAttribute{
				mBitType:         FieldAttr.FieldType,
				mLengthAttribute: FieldAttr.FieldLen,
				mLen:             len(FieldData),
				mData:            data,
				mIsSet:           true,
			}
		}
	case Bcd:
		var length int
		if FieldAttr.FieldLen == Fixed {
			length = FieldAttr.FieldMaxLength
		} else {
			length = len(FieldData)
		}
		if data, ok := hex.DecodeString(FieldData); ok == nil {
			i.mFieldsAttrs[FieldID] = bitAttribute{
				mBitType:         FieldAttr.FieldType,
				mLengthAttribute: FieldAttr.FieldLen,
				mLen:             length,
				mData:            data,
				mIsSet:           true,
			}
		}
	default:
		logging.GetLog().Infof("other types are not implemented %v", FieldAttr.FieldType)
	}
	i.mFieldEnabled = append(i.mFieldEnabled, FieldID)
	i.SetBit(FieldID)
}

//CheckBit check bit enable or not
func (i *Message) CheckBit(FieldID int) bool {
	var IsEnabled bool
	BitCheck := [8]byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	if i.mBitMap[(FieldID-1)/8]&BitCheck[(FieldID-1)%8] != 0 {
		IsEnabled = true
	} else {
		IsEnabled = false
	}

	return IsEnabled
}

//SetBit enabled bit
func (i *Message) SetBit(FieldID int) {
	BitCheck := [8]byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	i.mBitMap[(FieldID-1)/8] |= BitCheck[(FieldID-1)%8]
}

//ClearBit disabled bit
func (i *Message) ClearBit(FieldID int) {
	BitCheck := [8]byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	if IsEnabled := i.CheckBit(FieldID); IsEnabled {
		i.mBitMap[(FieldID-1)/8] ^= BitCheck[(FieldID-1)%8]
	}
}

//SetMTI set value MTI
func (i *Message) SetMTI(MTI int) {
	i.mMTI = MTI
}

//Pack make new a iso 8583 message format MTI+BITMAP+FIELD DATA
func (i *Message) Pack() ([]byte, int, error) {
	var err error
	var buffer = make([]byte, maxBufferSize)

	err = nil
	if MTI, ok := hex.DecodeString(BinToString(i.mMTI, 2)); ok == nil {
		copy(buffer, MTI)
		i.mPackageSize += 2
	} else {
		err = errors.New("MTI invalid type")
	}

	if i.mBitMap != nil {
		copy(buffer[i.mPackageSize:], i.mBitMap)
		i.mPackageSize += 8
	} else {
		err = errors.New("bitmap is nil")
	}

	sort.Ints(i.mFieldEnabled[:])
	//map in golang
	for _, field := range i.mFieldEnabled {
		for k, v := range i.mFieldsAttrs {
			if field == k {
				switch v.mLengthAttribute {
				case Fixed:
					copy(buffer[i.mPackageSize:], v.mData)
					i.mPackageSize += len(v.mData)
				case Llvar:
					if llvar, ok := hex.DecodeString(BinToString(v.mLen, 1)); ok == nil {
						copy(buffer[i.mPackageSize:], llvar)
						i.mPackageSize++
					}
					copy(buffer[i.mPackageSize:], v.mData)
					i.mPackageSize += (v.mLen + 1) / 2
				case Lllvar:
					if lllvar, ok := hex.DecodeString(BinToString(v.mLen, 2)); ok == nil {
						copy(buffer[i.mPackageSize:], lllvar)
						i.mPackageSize += 2
					}
					copy(buffer[i.mPackageSize:], v.mData)
					i.mPackageSize += (v.mLen + 1) / 2
				default:
					err = errors.New("length types are not implemented")
				}
			}
		}
	}
	i.mBuffer = make([]byte, i.mPackageSize)
	copy(i.mBuffer, buffer[:i.mPackageSize])
	return i.mBuffer, i.mPackageSize, err
}

//Unpack unpack iso8583 receive from client
func (i *Message) Unpack() {

	//skip 10 byte (start from field data)
	count := 10
	current := 10
	for field := 1; field <= maxBitmapSize*maxBitmapSize; field++ {
		if i.CheckBit(field) {
			if !Contains(i.mFieldEnabled, field) {
				i.mFieldEnabled = append(i.mFieldEnabled, field)
			}
			FieldAttr := Spec[field]

			var length int

			switch FieldAttr.FieldLen {
			case Fixed:
				length = FieldAttr.FieldMaxLength
			case Llvar:
				length, _ = HexToInt(i.mBuffer[count : count+1])
			case Lllvar:
				length, _ = HexToInt(i.mBuffer[count : count+2])
			default:
				logging.GetLog().Info("other types are not implemented")
			}
			switch FieldAttr.FieldType {
			case An, Ans:
				count += length
				i.mFieldsAttrs[field] = bitAttribute{
					mBitType:         FieldAttr.FieldType,
					mLengthAttribute: FieldAttr.FieldLen,
					mLen:             length,
					mData:            i.mBuffer[current:count],
					mIsSet:           true,
				}
				current = count
			case Bcd:
				count += (length + 1) / 2
				i.mFieldsAttrs[field] = bitAttribute{
					mBitType:         FieldAttr.FieldType,
					mLengthAttribute: FieldAttr.FieldLen,
					mLen:             length,
					mData:            i.mBuffer[current:count],
					mIsSet:           true,
				}
				current = count
			default:
				logging.GetLog().Info("other types are not implemented")
			}
		}
	}
	i.Parse()
}

// Parse print data for each field
func (i *Message) Parse() {
	var FieldID string
	if i.IsRequest() {
		logging.GetLog().Info("=============== Request Message ===============")
	} else {
		logging.GetLog().Info("=============== Response Message ===============")
	}
	logging.GetLog().Info("TPDU = ", hex.EncodeToString(i.mTPDU))
	logging.GetLog().Debug("Data = ", hex.EncodeToString(i.mBuffer))
	logging.GetLog().Info("Message Type = ", i.mMTI)
	for _, field := range i.mFieldEnabled {
		for k, v := range i.mFieldsAttrs {
			if field == k {
				if field < 10 {
					FieldID = fmt.Sprintf("00%d", field)
				} else if field < 99 {
					FieldID = fmt.Sprintf("0%d", field)
				} else if field < 999 {
					FieldID = fmt.Sprintf("%d", field)
				}
				logging.GetLog().Infoln("Field", FieldID, " = ", hex.EncodeToString(v.mData), " - ", Spec[field].FieldDescription)
			}
		}
	}
	logging.GetLog().Info("========================= End Parse Message =========================")
}

//IsRequest check msg is request or not
func (i *Message) IsRequest() bool {
	ret := ((i.mMTI / 10) % 2) == 0
	return ret
}

//SetResponseMTI set MTI from request MTI to response MTI
func (i *Message) SetResponseMTI() error {
	var err error
	err = nil
	if !i.IsRequest() {
		err = errors.New("This message are not a request")
		return err
	}

	if i.mMTI%2 == 1 {
		i.mMTI--
	}

	i.mMTI += 10
	return err
}

// IsLogon check the message is logon or not
func (i *Message) IsLogon() bool {
	return (((i.mMTI % 1000) / 800) == 1) && (((i.mMTI % 1000) % 800) < 100)
}

//IsReversalOrChargeBack  the check message is reversal or not
func (i *Message) IsReversalOrChargeBack() bool {
	return (((i.mMTI % 1000) / 400) == 1) && (((i.mMTI % 1000) % 400) < 100)
}

//IsFinAncial check the message is finalcial or not
func (i *Message) IsFinAncial() bool {
	return (((i.mMTI % 1000) / 200) == 1) && (((i.mMTI % 1000) % 200) < 100)
}

//IsAuthorization check the message is authorization or not
func (i *Message) IsAuthorization() bool {
	return (((i.mMTI % 1000) / 100) == 1) && (((i.mMTI % 1000) % 100) < 100)
}

//SwapNII swap dst nii And src nii
func (i *Message) SwapNII() {
	i.mTPDU[1], i.mTPDU[3] = i.mTPDU[3], i.mTPDU[1]
	i.mTPDU[2], i.mTPDU[4] = i.mTPDU[4], i.mTPDU[2]
}

//Clone make a new iso8583 data from exists iso8583 data
func (i *Message) Clone() *Message {
	m := &Message{
		mTPDU:         i.mTPDU,
		mMTI:          i.mMTI,
		mBitMap:       i.mBitMap,
		mBuffer:       i.mBuffer,
		mFieldEnabled: i.mFieldEnabled,
		mFieldsAttrs:  i.mFieldsAttrs,
	}
	return m
}

//BuildMsg build completed iso8583 msg 2 byte length + tpdu + bitmap + mti + field data
func (i *Message) BuildMsg() ([]byte, error) {
	var err error
	var msg = make([]byte, maxBufferSize)
	length, _ := hex.DecodeString(BinToString(len(i.mTPDU)+len(i.mBuffer), 2))

	copied := copy(msg, length)
	copied += copy(msg[copied:], i.mTPDU)
	copied += copy(msg[copied:], i.mBuffer)
	// length, _ := hex.DecodeString(BinToString(copied, 2))
	return msg[:copied], err
}
