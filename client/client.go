package main

import (
	"log"

	"github.com/drew138/graphics-app/src/messages"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := messages.NewImageServiceClient(conn)
}
