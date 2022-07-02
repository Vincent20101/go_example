//通讯协议处理，主要处理封包和解包的过程
package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	ConstHeader         = "www.01happy.com"
	ConstHeaderLength   = 15
	ConstSaveDataLength = 4
)

//封包
func Packet(message []byte) []byte {
	s := string(message)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(1 << 8)
	fmt.Println((int(s[0]) << 8) | int(s[1]))
	fmt.Println(string(s[0]))
	fmt.Println(string(s[1]))
	fmt.Println(string(s[2]))
	fmt.Println(string(s[3]))
	fmt.Println(string(s[4]))
	fmt.Println(string(s[5]))
	fmt.Println(string(s[6]))
	fmt.Println(string(s[7]))
	fmt.Println(string(s[8]))
	fmt.Println(string(s[42]))
	fmt.Println(string(s[43]))

	toBytes := IntToBytes(len(message))
	slice := []byte(ConstHeader)
	i := append(slice, toBytes...)
	i2 := append(i, message...)
	return i2
}

//解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+ConstHeaderLength+ConstSaveDataLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
			i2 := i + ConstHeaderLength + ConstSaveDataLength + messageLength
			if length < i2 {
				break
			}
			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
			t := i + ConstHeaderLength + ConstSaveDataLength + messageLength
			fmt.Println(t)
			readerChannel <- data

			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
