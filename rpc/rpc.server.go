package rpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type RPCServer struct{}

func (s *RPCServer) StartJob(ctx context.Context, in *JobRequest) (*JobReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &JobReply{Message: "Hello " + in.GetName()}, nil
}

func (s *RPCServer) StopJob(ctx context.Context, in *JobRequest) (*JobReply, error) {
	log.Printf("Received Agagin: %v", in.GetName)
	return &JobReply{Message: "Hello Agagin " + in.GetName()}, nil
}

func NewRPCServer() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen failed")
	}
	svr := grpc.NewServer()
	RegisterGreeterServer(svr, &RPCServer{})
	if err := svr.Serve(lis); err != nil {
		fmt.Printf("rpc failed: %v", err)
	}
}
