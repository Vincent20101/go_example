package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	// 十进制
	var a int = 10
	a2 := strconv.FormatInt(int64(a), 2)   //10进制转换为2进制
	fmt.Printf("%v \n", a2)                //1010
	a8 := strconv.FormatInt(int64(a), 8)   //10进制转换为8进制
	fmt.Printf("%v \n", a8)                //12
	a16 := strconv.FormatInt(int64(a), 16) //10进制转换为16进制
	fmt.Printf("%v \n", a16)               //a
	var b string = "1010"
	p, _ := strconv.ParseInt(b, 2, 64) // 2进制转10进制
	fmt.Printf("%v\n", p)              //10
	var b2 string = "345"
	p2, _ := strconv.ParseInt(b2, 8, 64) // 8进制转10进制
	fmt.Printf("%v\n", p2)               //229

	toBytes := IntToBytes(257)
	fmt.Println(toBytes)
	fmt.Println(BytesToInt(toBytes))

	var v3 uint32
	var b3 [4]byte
	v3 = 257
	b3[3] = uint8(v3)
	b3[2] = uint8(v3 >> 8)
	b3[1] = uint8(v3 >> 16)
	b3[0] = uint8(v3 >> 24)
	fmt.Println("b3:", b3)
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

// https://learnku.com/articles/61994
//很多编程语言都封装了 pack ()，unpack () 函数，用于把各种变量转为二进制。比如 php,python，ruby，go 也有 binary 包来帮我们做转化
//
//1.binary.Write ()，把变量写入到 [] byte
//
//2.binary.Read (), 读取 [] byte 转为对应类型的变量
//3. 需要了解位运算，与 |，左移 <<，右移>>
//
//比如封装了 [] byte 转 int16, int32 的函数
/*
以num = 300为例，计算机加载进去内存后，用两个字节保存 00000001 00101100(不同cup架构存放会不同，但是byte()会读取最低位的字节)
1.byte(),会把uint16读取一个低字节，所以byte(300)返回00101100，也就是44 属于uint16的低位
2.把uint16右移8位，得到00000000 00000001，再取一个低字节00000001,属于uint16的高位
大端序，字节数组从0开始放高位到地位，所以b[0] = 1, b[1]= 44
*/
func MyPutUint16(b []byte, num uint16) {
	_ = b[1]
	b[1] = byte(num)
	b[0] = byte(num >> 8)
}
func MyPutUint32(b []byte, num uint32) {
	_ = b[3]
	b[3] = byte(num)
	b[2] = byte(num >> 8)
	b[1] = byte(num >> 16)
	b[0] = byte(num >> 24)
}

// 封转了 int16,int32 转 [] byte
/*
以b[1,44]为例子
1.首先把每个字节的整数转uint16，表示用两个字节保存
00000000 00000001 //b[0]
00000000 00101100 //b[1]

2. 因为是大端序，需要把高字结位b[0]移动到左边。
00000001 00000000  //b[0]

3.使用与运算|, b[0]|b[1]得到新的uint16
00000001 00000000 b[0]
00000000 00101100 b[1]
00000001 00101100 uint16
*/
func MyUint16(b []byte) uint16 {
	_ = b[1]
	return uint16(b[1]) | uint16(b[0])<<8
}

//同理，可以得出uint32
func MyUint32(b []byte) uint32 {
	_ = b[3]
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}
