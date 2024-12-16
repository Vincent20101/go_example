package pfcpie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

const (
	// Casa enterprise Id
	CasaEnterpriseID     = 20858
	EricssonEnterpriseID = 193
)

const (
	Mask8 = 1<<8 - 1
	Mask7 = 1<<7 - 1
	Mask6 = 1<<6 - 1
	Mask5 = 1<<5 - 1
	Mask4 = 1<<4 - 1
	Mask3 = 1<<3 - 1
	Mask2 = 1<<2 - 1
	Mask1 = 1<<1 - 1
)

const (
	BitMask8 = 1 << 7
	BitMask7 = 1 << 6
	BitMask6 = 1 << 5
	BitMask5 = 1 << 4
	BitMask4 = 1 << 3
	BitMask3 = 1 << 2
	BitMask2 = 1 << 1
	BitMask1 = 1
)

func btou(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func utob(u uint8) bool {
	return u != 0
}

func ConvertStringToBCDCode(str string) []byte {
	var Val []byte
	strLen := len(str)
	if strLen == 0 {
		return Val
	}

	for index := 0; index < strLen; index += 2 {
		if index+1 >= strLen {
			break
		}
		num1, _ := strconv.Atoi(string(str[index]))
		num2, _ := strconv.Atoi(string(str[index+1]))
		num := ((byte(num2) << 4) & 0xf0) | (byte(num1) & 0x0f)

		Val = append(Val, num)
	}
	if strLen%2 != 0 {
		numTemp, _ := strconv.Atoi(string(str[strLen-1]))
		num := (byte(numTemp) & 0x0f) | 0xf0
		Val = append(Val, num)
	}

	return Val
}

func ConvertBCDCodeToString(buf []byte) string {
	var Str string
	if len(buf) == 0 {
		return Str
	}

	for _, data := range buf {
		str1 := strconv.Itoa(int(data & 0x0f))
		str2 := strconv.Itoa(int((data & 0xf0) >> 4))

		if len(str2) == 1 {
			Str = Str + str1 + str2
		} else {
			Str = Str + str1
		}

	}

	return Str
}

// Append uint16 to slice in BigEndian order and return a new slice.
func AppendUint16(s []byte, n uint16) []byte {
	t := make([]byte, 2)
	binary.BigEndian.PutUint16(t, n)
	return append(s, t...)
}

// Append uint32 to slice in BigEndian order and return a new slice.
func AppendUint32(s []byte, n uint32) []byte {
	t := make([]byte, 4)
	binary.BigEndian.PutUint32(t, n)
	return append(s, t...)
}

const JAN_1970 = 2208988800 // 1970 - 1900 in seconds

func TimeToTimestampValue(t time.Time) uint32 {
	return uint32(t.Unix()) + uint32(JAN_1970)
}

func TimestampValueToTime(ts uint32) time.Time {
	return time.Unix(int64(ts-uint32(JAN_1970)), 0).UTC()
}

type ByteReader struct {
	*bytes.Buffer
}

// NewByteReader return a new ByteReader.
func NewByteReader(b []byte) *ByteReader {
	return &ByteReader{bytes.NewBuffer(b)}
}

// IsDrained check if Reader buffer is drained.
func (b *ByteReader) IsDrained() bool {
	_, e := b.ReadByte()
	_ = b.UnreadByte()
	return e != nil
}

// ReadBytes read n bytes from ByteReader return byte slice.
// If the unread bytes is not enough, return left bytes and error.
func (b *ByteReader) ReadBytes(n int) (s []byte, err error) {
	// No unread bytes left in data.
	res := b.Next(n)
	l := len(res)
	if l < n {
		err = fmt.Errorf("inadequate bytes, require: %d, got: %d", s, l)
	}
	return res, err
}

// ReadUint16 reads Uint8 from Byte.
// If unread bytes is less than one, return 0, error
func (b *ByteReader) ReadUint8() (u uint8, err error) {
	t, e := b.ReadBytes(1)
	if e != nil {
		return 0, e
	}
	return uint8(t[0]), nil
}

//	ReadUint16 read Uint16 in BigEndian from Byte
//
// If unread bytes is less than 2, return 0, error
func (b *ByteReader) ReadUint16() (u uint16, err error) {
	t, e := b.ReadBytes(2)
	if e != nil {
		return 0, e
	}
	u = binary.BigEndian.Uint16(t)
	return u, nil
}

//	ReadUint32 read Uint32 in BigEndian from Byte
//
// If unread bytes is less than 4, return 0, error
func (b *ByteReader) ReadUint32() (u uint32, err error) {
	t, e := b.ReadBytes(4)
	if e != nil {
		return 0, e
	}
	u = binary.BigEndian.Uint32(t)
	return u, nil
}
