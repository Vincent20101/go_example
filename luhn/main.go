package main

import (
	"fmt"

	luhn "github.com/vincent20101/go-example/luhn/tool"
)

func main() {
	n, _ := luhn.Generate("49015420323751")
	fmt.Println(n)

	//左移
	var a int = 1
	a = a << 10
	fmt.Print("左移后的结果为:")
	fmt.Println(a)

	//右移
	var b int = 1024
	b = b >> 10
	fmt.Print("右移后的结果为:")
	fmt.Println(b)

	//fmt.Println(4 << 4)
	//fmt.Println(4 | 144)

	fmt.Println(148 >> 4)

	//byte_data := []byte{94}

	// 将 byte 装换为 16进制的字符串
	//hex_string_data := hex.EncodeToString(byte_data)
	// byte 转 16进制 的结果
	//println(hex_string_data)

	/* ====== 分割线 ====== */

	// 将 16进制的字符串 转换 byte
	//hexData, _ := hex.DecodeString("94")
	// 将 byte 转换 为字符串 输出结果
	//println(hexData)

}
