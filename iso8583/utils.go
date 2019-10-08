package iso8583

import (
	"encoding/hex"
	"errors"
	"fmt"
	"golang/logging"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//MessageLengthType alias type int for message type
type MessageLengthType int

// LengthBcd
const (
	LengthBcd MessageLengthType = iota
	LengthHex
)

// ToMessageLengthType convert string to MessageLengthType
func ToMessageLengthType(typ string) MessageLengthType {
	var result MessageLengthType
	switch typ {
	case "BCD":
		result = LengthBcd
	case "HEX":
		result = LengthHex
	default:
		logging.GetLog().Info("other types are not implemented")
	}
	return result
}

//ToHexString convert str to hex str
//ex: "01100191" -> "3031313030313931"
func ToHexString(str string) string {
	return fmt.Sprintf("%x", str)
}

//StringToAsc convert hex str to byte
//ex: "3031313030313931" -> []byte("01100191")
func StringToAsc(str string) ([]byte, error) {
	return hex.DecodeString(ToHexString(str))
}

//PadAmount padding amount with pad string
func PadAmount(str string, length int, pad string) string {
	if strings.Contains(str, ".") {
		str = strings.Replace(string(str), ".", "", -2)
	} else {
		str = str + "00"
	}
	return times(pad, length-len(str)) + str
}

//PadLeft padding left pad to str
func PadLeft(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

//PadRight padding right pad to str
func PadRight(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

//BinToString convert integer number to str
//ex: 920000 -> "920000"
func BinToString(Input, Length int) string {
	str := strconv.Itoa(Input)
	return PadLeft(str, Length*2, "0")
}

//FromHexChar convert byte hex to byte int
func FromHexChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}

	return 0, false
}

// MessageLengthToInt convert 2 byte len to int
func MessageLengthToInt(typ MessageLengthType, length []byte) (int, error) {
	var err error
	var msglen int64

	if len(length) > 2 {
		err = errors.New("bytes too long")
	}
	switch typ {
	case LengthHex:
		msglen, err = strconv.ParseInt(hex.EncodeToString(length[:2]), 16, 64)
	case LengthBcd:
		msglen, err = strconv.ParseInt(hex.EncodeToString(length[:2]), 16, 64)
	}
	return int(msglen), err
}

// HexToInt convert byte hex to int
func HexToInt(value []byte) (int, error) {
	v, err := strconv.ParseInt(hex.EncodeToString(value[:2]), 16, 64)
	return int(v), err
}

const (
	numericCharset      = "0123456789"
	alphaNumericCharset = "abcdefghijklmnopqrstuvwxyz0123456789" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//StringWithCharset return string with optional charset
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//NumericString return string with numberic charset
func NumericString(length int) string {
	return StringWithCharset(length, numericCharset)
}

//AlphaNumericString return string with alpha numeric charset
func AlphaNumericString(length int) string {
	return StringWithCharset(length, alphaNumericCharset)
}

// Contains tells whether a contains x.
func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
