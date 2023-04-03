package main

import (
	"context"
	"fmt"
	"net"

	"calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Calculator server unary func executed....")

	res := &calculatorpb.SumResponse{
		SumResult: req.GetFirstNumber() + req.GetLastNumber(),
	}

	return res, nil
}

func main() {
	fmt.Println("Calculator unary server started...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Printf("Failed to listen /n: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve /n: %v", err)
	}
}
