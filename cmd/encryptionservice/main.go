package main

import (
	"log"
	"net"

	"github.com/kordape/redbrain-challenge/internal/service"
	"github.com/kordape/redbrain-challenge/pkg/encryptorapi"
	"google.golang.org/grpc"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &service.EncryptionService{}

	// Register gRPC server
	encryptorapi.RegisterEncryptorServer(s, srv)

	// Listen on port 8080
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
