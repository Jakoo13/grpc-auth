package main

import (
	"log"
	"net"

	"github.com/Jakoo13/grpc-auth/calculator/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s\n", addr)

	serverInstance := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(serverInstance, &Server{})

	if err := serverInstance.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
