package main

import (
	"log"
	"net"

	"github.com/drew138/graphics-app/src/handlers"
	"github.com/drew138/graphics-app/src/messages"
	"google.golang.org/grpc"
)

func main() {
	port := "50051"
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Started listening on port: %v\n", port)
	s := grpc.NewServer()
	imageServer := &handlers.Server{}
	messages.RegisterImageServiceServer(s, imageServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
