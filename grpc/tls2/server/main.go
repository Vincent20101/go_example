package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"os"

	pb "github.com/vincent20101/go-example/grpc/tls2/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// Load server certificate and private key
	serverCert, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("failed to read server certificate: %v", err)
	}
	serverKey, err := os.ReadFile("server.key")
	if err != nil {
		log.Fatalf("failed to read server private key: %v", err)
	}

	// Create TLS credentials
	serverCreds := credentials.NewServerTLSFromCert(&tls.Certificate{
		Certificate: [][]byte{serverCert},
		PrivateKey:  serverKey,
	})
	//if err != nil {
	//	log.Fatalf("failed to create server credentials: %v", err)
	//}

	// Create gRPC server with TLS credentials
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(serverCreds))
	pb.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Start gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
