package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 创建一个字节数组
	data := []byte{1, 2, 3, 4, 5, 6}

	// 使用 NewBuffer 将字节数组包装成一个 Buffer
	buf := bytes.NewBuffer(data)

	// 读取 Buffer 中的数据
	// 创建一个大小为 2 的缓冲区来接收读取的数据
	readData := make([]byte, 2)
	n, err := buf.Read(readData)
	if err != nil {
		fmt.Println("Error reading from buffer:", err)
		return
	}
	fmt.Printf("Read %d bytes: %v\n", n, readData)

	// 获取 Buffer 中剩余数据的长度
	remainingLen := buf.Len()
	fmt.Println("Remaining data length in buffer:", uint32(remainingLen))

	//buf.Reset()
	//remainingLen = buf.Len()
	//fmt.Println("Remaining data length in buffer:", remainingLen)

	nextByte := buf.Next(4)
	fmt.Println("Next byte from buffer:", nextByte, buf.Len())
	capPktLen := binary.BigEndian.Uint32(nextByte)
	fmt.Println(capPktLen)
	// 读取剩余的数据
	// 创建一个大小为剩余数据长度的缓冲区来接收读取的数据
	remainingData := make([]byte, remainingLen)
	n, err = buf.Read(remainingData)
	if err != nil {
		fmt.Println("Error reading remaining data from buffer:", err)
		return
	}
	fmt.Printf("Read %d bytes remaining: %v\n", n, remainingData)

	// 使用 Next 方法获取下一个字节
	nextByte = buf.Next(1)
	fmt.Println("Next byte from buffer:", nextByte)
}
