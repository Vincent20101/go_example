package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/encoding"

	"github.com/vincent20101/go-example/grpc/gtest/test/codec"

	gproto "github.com/vincent20101/go-example/grpc/gtest/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *gproto.HelloRequest) (*gproto.HelloReply, error) {
	fmt.Println("test22")
	//time.Sleep(4 * time.Second)
	return &gproto.HelloReply{
		Message: "Hello service2 " + request.Name,
	}, nil
}

func main() {
	encoding.RegisterCodec(&codec.JSONCoder{})
	g := grpc.NewServer()
	//g := grpc.NewServer(grpc.CustomCodec(&codec.JSONCoder{}))
	s := Server{}
	gproto.RegisterGreeterServer(g, &s)
	listen, err := net.Listen("tcp4", ":8081")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(listen)
}
