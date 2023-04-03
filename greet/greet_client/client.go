package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"grpc-udamy/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect /n: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	// unary(c)
	doServerStreaming(c)

}

func unary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		FirstName: "Mukim ",
		LastName:  "Billah",
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		fmt.Printf("Failed to call greet func /n: %v", err)
	}

	fmt.Printf("Response form greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greet: &greetpb.GreetRequest{
			FirstName: "Mukim ",
			LastName:  "Billah",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		fmt.Printf("Failed to call greet func /n: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read stream /n: %v", err)
		}

		fmt.Println("Response from GreetManyTimes: ", msg.Result)
	}
}
