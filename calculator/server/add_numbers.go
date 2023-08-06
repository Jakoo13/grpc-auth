package main

import (
	"context"
	"log"

	"github.com/Jakoo13/grpc-auth/calculator/proto"
)

func (s *Server) AddNumbers(ctx context.Context, in *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	log.Printf("Add Numbers function was invoked with %v\n", in)

	addedNumbers := in.FirstNumber + in.SecondNumber
	return &proto.CalculatorResponse{
		Result: addedNumbers,
	}, nil
}
