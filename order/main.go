package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/hoangduc02011998/k8s-go-grpc-order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAddServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(cxt context.Context, r *pb.AddRequest) (*pb.AddResponse, error) {
	result := &pb.AddResponse{}
	result.Result = r.A + r.B

	logMessage := fmt.Sprintf("A: %d   B: %d     sum: %d", r.A, r.B, result.Result)
	log.Println(logMessage)

	return result, nil
}
