package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"github.com/vincent20101/go-example/grpc/grpc_token_auth_test/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求")
		md, ok := metadata.FromIncomingContext(ctx)
		fmt.Println(md)
		if !ok {
			//已经开始接触到grpc的错误处理了
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		var (
			appid  string
			appkey string
		)

		if va1, ok := md["appid"]; ok {
			appid = va1[0]
		}

		if va1, ok := md["appkey"]; ok {
			appkey = va1[0]
		}

		if appid != "101010" || appkey != "i am key" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		res, err := handler(ctx, req)
		fmt.Println("请求已经完成")
		return res, err
	}

	opt := grpc.UnaryInterceptor(interceptor)
	var opts []grpc.ServerOption
	opts = append(opts, opt)

	tls, err := credentials.NewServerTLSFromFile("/root/project/goex/go_example/grpc/tls/cert/test/server.crt", "/root/project/goex/go_example/grpc/tls/cert/test/server.key")
	if nil != err {
		log.Fatalf("failed to create TLS: %v", err)
	}
	opts = append(opts, grpc.Creds(tls))
	opts = append(opts, grpc.InitialWindowSize(65535))
	opts = append(opts, grpc.MaxConcurrentStreams(65535))

	g := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}