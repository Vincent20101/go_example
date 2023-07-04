package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	pb "github.com/vincent20101/go-example/grpc/tls2/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

func main() {
	// Load server certificate
	serverCert, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("failed to read server certificate: %v", err)
	}

	// Create certificate pool
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(serverCert) {
		log.Fatalf("failed to append server certificate to certificate pool")
	}

	// Create TLS credentials
	clientCreds := credentials.NewTLS(&tls.Config{
		ServerName: "localhost",
		RootCAs:    certPool,
	})

	// Create gRPC connection with TLS credentials
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	c := pb.NewGreeterClient(conn)

	// Send request
	name := "world"
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s\n", r.Message)
}
