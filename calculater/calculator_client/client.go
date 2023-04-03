package main

import (
	"context"
	"fmt"

	"calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect /n: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	unary(c)

}

func unary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber: 20,
		LastNumber:  10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		fmt.Printf("Failed to call greet func /n: %v", err)
	}

	fmt.Printf("Response form calculator: %v", res.SumResult)
}
