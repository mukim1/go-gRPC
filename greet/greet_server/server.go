package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"grpc-udamy/greetpb"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet server unary func executed....")
	fitst_name := req.GetFirstName()
	last_name := req.GetLastName()

	res := &greetpb.GreetResponse{
		Result: "Hello " + fitst_name + last_name,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("Greet server streaming func executed....")
	first_name := req.Greet.GetFirstName()

	for i := 0; i < 30; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: "Hello " + first_name + " number " + fmt.Sprint(i),
		}
		stream.Send(res)
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func main() {
	fmt.Println("Greet server started...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Printf("Failed to listen /n: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve /n: %v", err)
	}
}
