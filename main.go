package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"

	imageService "github.com/drew138/graphics-app/src/image-service"
	"github.com/drew138/graphics-app/src/messages"
	"golang.org/x/crypto/acme/autocert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := "50051"

	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Started listening on port: %v\n", port)

	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("tls"),
		HostPolicy: autocert.HostWhitelist("localhost"),
	}
	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	creds := credentials.NewTLS(&tls.Config{GetCertificate: certManager.GetCertificate})
	server := grpc.NewServer(grpc.Creds(creds))

	reflection.Register(server)

	imageServiceServer := &imageService.ImageServiceServer{}
	messages.RegisterImageServiceServer(server, imageServiceServer)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
