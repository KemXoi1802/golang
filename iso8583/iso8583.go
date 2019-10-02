package iso8583

import (
	"encoding/hex"
	"fmt"
	"golang/utils"
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
func NewIso8583Data(data []byte) *ISO8583Data {
	return &ISO8583Data{
		mBuffer: data[5:],
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
	case BINARY:
		if data, ok := hex.DecodeString(FieldData); ok == nil {
			i.mFieldsAttrs[FieldID] = bitAttribute{
				mBitType:         FieldAttr.FieldType,
				mLengthAttribute: FieldAttr.FieldLen,
				mLen:             (len(FieldData) + 1) / 2,
				mData:            data,
				mIsSet:           true,
			}
		}
	default:
		fmt.Println("other types are not implemented")
		utils.GetLog().Infof("other types are not implemented %v", FieldAttr.FieldType)
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

// Parse print for each field
func (i *ISO8583Data) Parse() {
	var FieldID string
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
				utils.GetLog().Infoln(FieldID, " - ", hex.EncodeToString(v.mData))
			}
		}
	}
}

// type Data struct {
// 	string
// 	original interface{}
// }

// func (d Data) String() string {
// 	return d.string
// }

// func (d Data) Number() (int64, error) {
// 	return strconv.ParseInt(d.string, 10, 64)
// }

// func (d Data) Binary() (r string, err error) {
// 	hex, err := strconv.ParseInt(d.string, 16, 64)
// 	if err != nil {
// 		return
// 	}
// 	r = strconv.FormatInt(hex, 2)
// 	return
// }

// func Parse(raw string) (msg Message, err error) {
// 	msg.Data = map[int]Data{}
// 	msg.MTI = raw[0:4]
// 	raw = raw[4:]

// 	if raw, err = msg.parseBitMap(raw); err != nil {
// 		return
// 	}

// 	if err = msg.parseData(raw); err != nil {
// 		return
// 	}

// 	return
// }

// func (m *Message) parseBitMap(raw string) (r string, err error) {
// 	if m.BitMap, err = toBitMap(raw[0:16]); err != nil {
// 		return
// 	}
// 	m.RawBitMap = raw[0:16]
// 	raw = raw[16:]

// 	if m.BitMap[0] == '1' {
// 		var secondary string
// 		if secondary, err = toBitMap(raw[0:16]); err != nil {
// 			return
// 		}
// 		m.BitMap += secondary
// 		m.RawBitMap += raw[0:16]
// 		raw = raw[16:]
// 	}

// 	for i, _ := range m.BitMap {
// 		b, _ := strconv.Atoi(string(m.BitMap[i]))
// 		if i > 0 && b == 1 {
// 			m.Fields = append(m.Fields, i+1)
// 		}
// 	}

// 	r = raw

// 	return
// }

// func toBitMap(src string) (dst string, err error) {
// 	for _, code := range src {
// 		n, err := strconv.ParseInt(string(code), 16, 64)
// 		if err != nil {
// 			return "", fmt.Errorf("strconv.ParseInt(string(%s), 16, 64): %s", code, err)
// 		}
// 		code := strconv.FormatInt(n, 2)
// 		for len(code) < 4 {
// 			code = "0" + code
// 		}
// 		dst += code
// 	}

// 	for len(dst) < 64 {
// 		dst = "0" + dst
// 	}

// 	return
// }

// func (m *Message) parseData(raw string) (err error) {
// 	for _, field := range m.Fields {
// 		elem := Spec[field]
// 		var data string
// 		if elem.FieldType == BINARY {
// 			data = raw[:elem.FieldMaxLength/4]
// 			raw = raw[elem.FieldMaxLength/4:]
// 			goto setData
// 		}

// 		if elem.FieldLen == FIXED {
// 			data = raw[:elem.FieldMaxLength]
// 			raw = raw[elem.FieldMaxLength:]
// 			goto setData
// 		}

// 		{
// 			var elen int
// 			if elem.FieldLen == LLVAR {
// 				elen = 2
// 			} else if elem.FieldLen == LLLVAR {
// 				elen = 3
// 			}

// 			var l int64
// 			if l, err = strconv.ParseInt(raw[:elen], 10, 64); err != nil {
// 				return
// 			}
// 			raw = raw[elen:]
// 			data = raw[:l]
// 			raw = raw[l:]
// 		}

// 	setData:
// 		m.Data[field] = Data{string: data}
// 	}
// 	return
// }

// func (iso *ISO8583Data) PackField(FieldId int, FieldData string) error {
// 	var err error
// 	return err
// }

// func (m *Message) AddData(field int, value interface{}) (err error) {
// 	if m.Data == nil {
// 		m.Data = map[int]Data{}
// 	}
// 	data := Data{original: value}
// 	switch Spec[field].FieldType {
// 	case BCD:
// 		if i, ok := value.(int); ok {
// 			data.string = strconv.Itoa(i)
// 		} else {
// 			return fmt.Errorf("value type is %T; want int", value)
// 		}
// 	case BINARY:
// 		var binary int64
// 		if s, ok := value.(string); ok {
// 			if binary, err = strconv.ParseInt(s, 2, 64); err != nil {
// 				return
// 			}
// 		} else {
// 			return fmt.Errorf("value type is %T; want string", value)
// 		}
// 		data.string = strconv.FormatInt(binary, 16)
// 	default:
// 		if s, ok := value.(string); ok {
// 			data.string = s
// 		} else {
// 			return fmt.Errorf("value type is %T; want string", value)
// 		}
// 	}
// 	m.Data[field] = data
// 	return
// }

// func (m *Message) Build() (r string) {
// 	m.genBitMap()
// 	r += m.MTI
// 	r += m.RawBitMap
// 	r += m.packDataElem()
// 	return
// }

// func (m *Message) genBitMap() {
// 	var bitmap [128]int
// 	m.Fields = []int{}
// 	for field, _ := range m.Data {
// 		m.Fields = append(m.Fields, field)
// 		bitmap[field-1] = 1
// 		if field > 64 {
// 			bitmap[0] = 1
// 		}
// 	}
// 	sort.Sort(sort.IntSlice(m.Fields))
// 	m.BitMap = ""
// 	for i, bit := range bitmap {
// 		if i > 64 && bitmap[0] != 1 {
// 			break
// 		}
// 		m.BitMap += strconv.Itoa(bit)
// 	}
// 	m.RawBitMap = ""
// 	for i := 1; i <= len(m.BitMap)/4; i++ {
// 		data, _ := strconv.ParseInt(m.BitMap[(i-1)*4:i*4], 2, 64)
// 		m.RawBitMap += strconv.FormatInt(data, 16)
// 	}
// }

// func (m *Message) packDataElem() (r string) {
// 	for _, field := range m.Fields {
// 		elem := Spec[field]
// 		data := m.Data[field]
// 		if elem.FieldLen == FIXED {
// 			switch elem.FieldType {
// 			case BCD, BINARY:
// 				r += fmt.Sprintf("%0"+fmt.Sprint(elem.FieldMaxLength)+"s", data.string)
// 				// println(field, fmt.Sprintf("%0"+fmt.Sprint(elem.len)+"s", data.string))
// 			// case "b":
// 			default:
// 				r += fmt.Sprintf("% "+fmt.Sprint(elem.FieldMaxLength)+"s", data.string)
// 				// println(field, fmt.Sprintf("% "+fmt.Sprint(elem.len)+"s", data.string))
// 			}
// 		} else if elem.FieldLen == LLVAR {
// 			elen := "2"
// 			r += fmt.Sprintf("%0"+elen+"d", len(data.string)) + data.string
// 			// println(field, fmt.Sprintf("%0"+elen+"d", len(data.string))+data.string)
// 		} else if elem.FieldLen == LLLVAR {
// 			elen := "3"
// 			r += fmt.Sprintf("%0"+elen+"d", len(data.string)) + data.string
// 			// println(field, fmt.Sprintf("%0"+elen+"d", len(data.string))+data.string)
// 		}
// 	}
// 	return
// }
// func LogISOMSG(data Message) {
// 	k := 0
// 	for ; k <= 128; k++ {
// 		for _, f := range data.Fields {
// 			if k == f {
// 				var FieldId string
// 				if f < 10 {
// 					FieldId = fmt.Sprintf("00%d", f)
// 				} else if f < 99 {
// 					FieldId = fmt.Sprintf("0%d", f)
// 				} else if f < 999 {
// 					FieldId = fmt.Sprintf("%d", f)
// 				}
// 				fmt.Println(FieldId, " - ", data.Data[f], Spec[f].FieldDescription)
// 			}

// 		}
// 	}
// }
