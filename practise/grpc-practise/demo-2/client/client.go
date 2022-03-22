package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	runtimeapiV1alpha2 "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"unsafe"

	"log"
	"time"
)

var (
	addr = "unix:///var/run/dockershim.sock"
)


func main() {

	filter := &runtimeapi.ContainerFilter{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect rpc server %v", err)
	}
	defer conn.Close()
	runtimeClient := runtimeapi.NewRuntimeServiceClient(conn)
	resp, err := runtimeClient.ListContainers(ctx, &runtimeapiV1alpha2.ListContainersRequest{
		Filter: v1alpha2ContainerFilter(filter),
	})
	if err != nil {
		log.Fatalf("failed to sayhello %v", err)
	}
	log.Printf("say hello %v", resp)
}

func v1alpha2ContainerFilter(from *runtimeapi.ContainerFilter) *v1alpha2.ContainerFilter {
	return (*v1alpha2.ContainerFilter)(unsafe.Pointer(from))
}
