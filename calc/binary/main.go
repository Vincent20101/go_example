package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
)

// https://www.jianshu.com/p/ec461b39bf43
func main() {
	//test1()
	//test2()
	test3()

}

//参数列表：
//1）r  可以读出字节流的数据源
//2）order  特殊字节序，包中提供大端字节序和小端字节序
//3）data  需要解码成的数据
//返回值：error  返回错误
//功能说明：Read从r中读出字节数据并反序列化成结构数据。data必须是固定长的数据值或固定长数据的slice。从r中读出的数据可以使用特殊的 字节序来解码，并顺序写入value的字段。当填充结构体时，使用(_)名的字段讲被跳过。
func test1() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi) // 3.141592653589793
}

//参数列表：
//1）w  可写入字节流的数据
//2）order  特殊字节序，包中提供大端字节序和小端字节序
//3）data  需要解码的数据
//返回值：error  返回错误
//功能说明：
//Write讲data序列化成字节流写入w中。data必须是固定长度的数据值或固定长数据的slice，或指向此类数据的指针。写入w的字节流可用特殊的字节序来编码。另外，结构体中的(_)名的字段讲忽略。
func test2() {
	buf := new(bytes.Buffer)
	pi := math.Pi

	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.Bytes())
}

//参数列表：v  需要计算长度的数据
//返回值：int 数据序列化之后的字节长度
//功能说明：
//Size讲返回数据系列化之后的字节长度，数据必须是固定长数据类型、slice和结构体及其指针。
func test3() {
	var a int
	p := &a
	b := [10]int64{1}
	s := "adsa"
	bs := make([]byte, 10)

	fmt.Println(binary.Size(a))  // -1
	fmt.Println(binary.Size(p))  // -1
	fmt.Println(binary.Size(b))  // 80
	fmt.Println(binary.Size(s))  // -1
	fmt.Println(binary.Size(bs)) // 10
}
