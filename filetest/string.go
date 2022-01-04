package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s1 := "字符串"
	s2 := "拼接"

	// 第一种
	s3 := s1 + s2
	fmt.Println(s3) //s3 = "打印字符串"

	// 第二种
	s33 := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(s33) //s33 = "打印字符串"

	// 第三种
	//需要先导入strings包
	//定义一个字符串数组包含上述的字符串
	var str []string = []string{s1, s2}
	//调用Join函数
	s34 := strings.Join(str, "")
	fmt.Println(s34)

	// 第四种
	//需要先导入bytes包
	//定义Buffer类型
	var bt bytes.Buffer
	//向bt中写入字符串
	bt.WriteString(s1)
	bt.WriteString(s2)
	//获得拼接后的字符串
	s35 := bt.String()
	fmt.Println(s35)

	// 第五种
	var build strings.Builder
	build.WriteString(s1)
	//build.Reset()
	build.WriteString(s2)
	s36 := build.String()
	fmt.Println(s36)

}
