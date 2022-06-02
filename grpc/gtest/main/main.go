package main

import (
	"encoding/json"
	"fmt"

	gproto "github.com/vincent20101/go-example/grpc/gtest/proto"

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
	fmt.Println(req.ProtoReflect().Descriptor())
	fmt.Println(req.ProtoReflect())
	fmt.Println(gproto.File_grpc_proto_helloworld_proto)
}
