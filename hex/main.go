package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(1 & 1)
	fmt.Println(string([]byte{1, 2, 3, 4, 5}))

	fmt.Println(strconv.FormatInt(65535, 2))
	fmt.Println(strconv.ParseInt("1111111111111111", 2, 64))
	fmt.Println(65535 << 8)
	fmt.Println(string([]byte{1, 1, 1, 1, 1}))
	//panic("aaa")
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

	v := "31323334353637383132333435363031"
	fmt.Println(len(v))
	//base64.StdEncoding.EncodeToString()
	//base64.URLEncoding.EncodeToString()

	p, _ := strconv.ParseInt(v, 16, 64) // 2进制转10进制
	fmt.Printf("%v\n", p)

	ss := "myself test"
	fmt.Printf("%#x\n", ss)
	fmt.Printf("%#X\n", ss)

	fmt.Printf("%x\n", 13)

}
