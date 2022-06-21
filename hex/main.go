package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	s := "测试数据"
	s2 := hex.EncodeToString([]byte(s))
	b, err := hex.DecodeString(s2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
	fmt.Println(s2)

	// fmt 包的格式化也可以将字节转为十六进制字符串
	s3 := fmt.Sprintf("%x", b)
	fmt.Println(s3)

	//byteData := []byte(`测试数据`)
	//
	//// 将 byte 装换为 16进制的字符串
	//hexStringData := hex.EncodeToString(byteData)
	//// byte 转 16进制 的结果
	//println(hexStringData)
	//
	///* ====== 分割线 ====== */
	//
	//// 将 16进制的字符串 转换 byte
	//hexData, _ := hex.DecodeString(hexStringData)
	//// 将 byte 转换 为字符串 输出结果
	//println(string(hexData))

}
