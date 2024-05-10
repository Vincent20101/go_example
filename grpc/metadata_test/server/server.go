package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/vincent20101/go-example/grpc/metadata_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	time.Sleep(3 * time.Second)
	fmt.Println("a的值:", ctx.Value("a"))
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	if !ok {
		fmt.Println("get metadata error")
	}
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
