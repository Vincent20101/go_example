package main

import (
	"encoding/base64"
	"fmt"

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

}
