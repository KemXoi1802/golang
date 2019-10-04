package iso8583

import (
	"encoding/hex"
	"fmt"
	"golang/logging"
	"strconv"
)

//bitAttribute
type bitAttribute struct {
	mBitType         BitType
	mLengthAttribute BitLength
	mMaxLen          int
	mLen             int
	mData            []byte
	mIsSet           bool
}

//ISO8583Data
type ISO8583Data struct {
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
func DefaultIso8583Data() *ISO8583Data {
	return &ISO8583Data{
		mTPDU:        make([]byte, maxTpduSize),
		mBitMap:      make([]byte, maxBitmapSize),
		mFieldsAttrs: make(map[int]bitAttribute),
		mBuffer:      make([]byte, maxBufferSize),
	}
}

//NewIso8583Data init
func NewIso8583Data(data []byte, size int) *ISO8583Data {
	var mti int
	var err error
	if mti, err = strconv.Atoi(hex.EncodeToString(data[5:7])); err != nil {
		logging.GetLog().Info("MTI is not type int")
	}
	// if len(data) != size {
	// 	logging.GetLog().Info("DATA or LENGTH invalid")
	// }
	return &ISO8583Data{
		//tpdu, mti, bitmap.
		mTPDU:        data[:5],
		mMTI:         mti,
		mBitMap:      data[7:15],
		mBuffer:      data[15:],
		mFieldsAttrs: make(map[int]bitAttribute),
		mPackageSize: size,
	}
}

//PackField pack data into field
func (i *ISO8583Data) PackField(FieldID int, FieldData string) {
	FieldAttr := Spec[FieldID]

	switch FieldAttr.FieldType {
	case AN, ANS:
		if data, ok := StringToAsc(FieldData); ok == nil {
			i.mFieldsAttrs[FieldID] = bitAttribute{
				mBitType:         FieldAttr.FieldType,
				mLengthAttribute: FieldAttr.FieldLen,
				mLen:             len(FieldData),
				mData:            data,
				mIsSet:           true,
			}
		}
	case BCD:
		var length int
		if FieldAttr.FieldLen == FIXED {
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
func (i *ISO8583Data) CheckBit(FieldID int) bool {
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
func (i *ISO8583Data) SetBit(FieldID int) {
	BitCheck := [8]byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	i.mBitMap[(FieldID-1)/8] |= BitCheck[(FieldID-1)%8]
}

//ClearBit disabled bit
func (i *ISO8583Data) ClearBit(FieldID int) {
	BitCheck := [8]byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	if IsEnabled := i.CheckBit(FieldID); IsEnabled {
		i.mBitMap[(FieldID-1)/8] ^= BitCheck[(FieldID-1)%8]
	}
}

//SetMTI
func (i *ISO8583Data) SetMTI(MTI int) {
	i.mMTI = MTI
}

//Pack
func (i *ISO8583Data) Pack() ([]byte, int, error) {
	var err error

	if MTI, ok := hex.DecodeString(BinToString(i.mMTI, 2)); ok == nil {
		copy(i.mBuffer, MTI)
		i.mPackageSize += 2
		fmt.Errorf("value type is %T; want int", i.mMTI)
	}

	if i.mBitMap != nil {
		copy(i.mBuffer[i.mPackageSize:], i.mBitMap)
		i.mPackageSize += 8
		fmt.Errorf("value type is %T; want byte array", i.mBitMap)
	}

	//map in golang
	for _, field := range i.mFieldEnabled {
		for k, v := range i.mFieldsAttrs {
			if field == k {
				switch v.mLengthAttribute {
				case FIXED:
					copy(i.mBuffer[i.mPackageSize:], v.mData)
					i.mPackageSize += len(v.mData)
				case LLVAR:
					if llvar, ok := hex.DecodeString(BinToString(v.mLen, 1)); ok == nil {
						copy(i.mBuffer[i.mPackageSize:], llvar)
						i.mPackageSize++
					}
					copy(i.mBuffer[i.mPackageSize:], v.mData)
					i.mPackageSize += (v.mLen + 1) / 2
				case LLLVAR:
					if lllvar, ok := hex.DecodeString(BinToString(v.mLen, 2)); ok == nil {
						copy(i.mBuffer[i.mPackageSize:], lllvar)
						i.mPackageSize += 2
					}
					copy(i.mBuffer[i.mPackageSize:], v.mData)
					i.mPackageSize += (v.mLen + 1) / 2
				default:
					fmt.Errorf("length type is not support")
				}
			}
		}
	}

	return i.mBuffer, i.mPackageSize, err
}

//Unpack unpack iso8583 receive from client
func (i *ISO8583Data) Unpack() {
	count := 0
	current := 0
	for field := 1; field <= maxBitmapSize*maxBitmapSize; field++ {
		if i.CheckBit(field) {
			i.mFieldEnabled = append(i.mFieldEnabled, field)
			FieldAttr := Spec[field]

			var length int

			switch FieldAttr.FieldLen {
			case FIXED:
				length = FieldAttr.FieldMaxLength
			case LLVAR:
				length, _ = HexToInt(i.mBuffer[count : count+1])
			case LLLVAR:
				length, _ = HexToInt(i.mBuffer[count : count+2])
			default:
				logging.GetLog().Info("other types are not implemented")
			}
			switch FieldAttr.FieldType {
			case AN, ANS:
				count += length
				i.mFieldsAttrs[field] = bitAttribute{
					mBitType:         FieldAttr.FieldType,
					mLengthAttribute: FieldAttr.FieldLen,
					mLen:             length,
					mData:            i.mBuffer[current:count],
					mIsSet:           true,
				}
				current = count
			case BCD:
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
func (i *ISO8583Data) Parse() {
	var FieldID string
	logging.GetLog().Info("=============== Full Message ===============")
	logging.GetLog().Info("TPDU = ", hex.EncodeToString(i.mTPDU))
	logging.GetLog().Info("Bit Map = ", hex.EncodeToString(i.mBitMap))
	logging.GetLog().Info("Data = ", hex.EncodeToString(i.mBuffer))
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
	logging.GetLog().Info("=============== End Full Message ===============")
}
