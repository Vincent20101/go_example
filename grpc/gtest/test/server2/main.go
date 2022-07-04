package main

import (
	"context"
	"fmt"
	"net"
	"time"

	gproto "github.com/vincent20101/go-example/grpc/gtest/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *gproto.HelloRequest) (*gproto.HelloReply, error) {
	time.Sleep(10 * time.Second)
	fmt.Println("test22")
	return &gproto.HelloReply{
		Message: "Hello service2 " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	s := Server{}
	gproto.RegisterGreeterServer(g, &s)
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(listen)
}
