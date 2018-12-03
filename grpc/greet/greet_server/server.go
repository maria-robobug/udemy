package main

import (
	"fmt"
	"log"
	"net"

	"github.com/maria-robobug/udemy/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello World!")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
