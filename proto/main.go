package main

import (
	"encoding/base64"
	"fmt"
	"regexp"

	gproto "github.com/vincent20101/go-example/grpc/gtest/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	//s := struct {
	//	name string
	//	age  int
	//}{
	//	name: "hangzai",
	//	age:  12,
	//}
	s := gproto.HelloRequest{Name: "hangzai"}
	b, err := proto.Marshal(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(base64.StdEncoding.EncodeToString(b))

	input := "My email is xxx@gmail.com xxx@abc.com yyy@bcd.com.cn"
	exp, _ := regexp.Compile("([a-zA-Z0-9]+)@[a-zA-Z0-9]+.[a-zA-Z0-9]+")
	fmt.Println(exp.FindString(input))
	fmt.Println(exp.FindAllString(input, -1))

	input = "sdfsdfsdfsdfsdf"
	exp2 := regexp.MustCompile("^[?!^/]*$")
	fmt.Println(exp2.FindAllString(input, -1))
}
