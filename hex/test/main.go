package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"unsafe"
)

func main() {
	byteData := []byte(`测试数据`)

	// 将 byte 装换为 16进制的字符串
	hexStringData := hex.EncodeToString(byteData)
	// byte 转 16进制 的结果
	println(hexStringData)

	/* ====== 分割线 ====== */

	// 将 16进制的字符串 转换 byte
	hexData, _ := hex.DecodeString(hexStringData)
	// 将 byte 转换 为字符串 输出结果
	println(string(hexData))

	input := "=小信1n!"
	fmt.Println("input:", input)

	e := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println("StdEncoding:", e, url.QueryEscape(e))
	// StdEncoding: PeWwj+S/oTFuIQ== PeWwj%2BS%2FoTFuIQ%3D%3D

	decodeString, err := base64.StdEncoding.DecodeString(e)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(decodeString))
	e = base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Println("URLEncoding:", e, url.QueryEscape(e))
	// URLEncoding: PeWwj-S_oTFuIQ== PeWwj-S_oTFuIQ%3D%3D

	e = base64.RawStdEncoding.EncodeToString([]byte(input))
	fmt.Println("RawStdEncoding:", e, url.QueryEscape(e))
	// RawStdEncoding: PeWwj+S/oTFuIQ PeWwj%2BS%2FoTFuIQ

	e = base64.RawURLEncoding.EncodeToString([]byte(input))
	fmt.Println("RawURLEncoding:", e, url.QueryEscape(e))
	// RawURLEncoding: PeWwj-S_oTFuIQ PeWwj-S_oTFuIQ

	s := "测试"
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(234 | 198)
}
