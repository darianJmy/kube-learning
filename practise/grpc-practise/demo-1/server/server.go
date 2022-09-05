package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "kube-learning/practise/grpc-practise/demo-1/alice"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedAliceServer
}

func (s *server) GetAlice(ctx context.Context, in *pb.AliceRequest) (*pb.AliceReply, error) {
	return &pb.AliceReply{Message: fmt.Sprintf("%s %d", in.GetName(), in.GetId())}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":30000")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAliceServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
