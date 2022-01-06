package main

import (
	"context"
	"fmt"
	"log"

	"GRPC/unary/sumpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	enableTLS := true
	opts := grpc.WithInsecure()

	if enableTLS {
		fmt.Println("TLS enabled")
		creds, err := credentials.NewClientTLSFromFile("cert/ca-cert.pem", "")
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		opts = grpc.WithTransportCredentials(creds)
	} else {
		fmt.Println("TLS disabled")
	}

	conn, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := sumpb.NewSumClient(conn)

	// numbers to add
	num := sumpb.Numbers{
		A: 15,
		B: 5,
	}

	// call Add service
	res, err := c.Add(context.Background(), &sumpb.SumRequest{Numbers: &num})
	if err != nil {
		log.Fatalf("failed to call Add: %v", err)
	}
	fmt.Println(res.Result)
}
