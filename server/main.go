package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"GRPC/unary/sumpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func main() {
	//enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	enableTLS := true
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	// Create the TLS credentials
	if enableTLS {
		fmt.Println("TLS enabled")
		creds, err := credentials.NewServerTLSFromFile("cert/server-cert.pem", "cert/server-key.pem")
		if err != nil {
			log.Fatalf("could not load TLS keys: %s", err)
		}
		opts = append(opts, grpc.Creds(creds))
	} else {
		fmt.Println("TLS disabled")
	}
	s := grpc.NewServer(opts...)
	sumpb.RegisterSumServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}

// Add returns sum of two integers
func (*server) Add(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	a, b := req.GetNumbers().GetA(), req.GetNumbers().GetB()
	sum := a + b
	return &sumpb.SumResponse{Result: sum}, nil
}
