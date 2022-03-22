package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "kube-learning/practise/grpc-practise/demo-1/alice"
	"log"
	"time"
)

var (
	addr = "127.0.0.1:30000"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect rpc server %v", err)
	}
	defer conn.Close()
	c := pb.NewAliceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetAlice(ctx, &pb.AliceRequest{Id: 12345, Name: "jixingxing"})
	if err != nil {
		log.Fatalf("failed to sayhello %v", err)
	}
	log.Printf("say hello %v", r.Message)
}
