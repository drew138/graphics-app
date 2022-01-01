package main

import (
	"log"

	"github.com/drew138/graphics-app/src/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := messages.NewImageServiceClient(conn)
	// c.ApplyBoxBlurFilter(context.Background())
}
