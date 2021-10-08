package main

import (
	"encoding/json"
	"fmt"
	gproto "go-example/grpc/proto"

	"github.com/golang/protobuf/proto"
)

func main() {

	req := gproto.HelloRequest{
		Name: "bobby",
	}
	s, _ := proto.Marshal(&req)
	fmt.Println(len(s))
	m, _ := json.Marshal(&req)
	fmt.Println(len(m))
	newReq := gproto.HelloRequest{}
	proto.Unmarshal(s, &newReq)
	fmt.Println(newReq.Name)
}
