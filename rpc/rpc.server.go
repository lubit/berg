package rpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

var (
	rpcServer *grpc.Server
	rpcOnce   sync.Once
)

type RPCService struct{}

func (s *RPCService) StartJob(ctx context.Context, in *JobRequest) (*JobReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &JobReply{Message: "Hello " + in.GetName()}, nil
}

func (s *RPCService) StopJob(ctx context.Context, in *JobRequest) (*JobReply, error) {
	log.Printf("Received Agagin: %v", in.GetName)
	return &JobReply{Message: "Hello Agagin " + in.GetName()}, nil
}

func StartRPCServer() error {
	rpcOnce.Do(func() {
		lis, err := net.Listen("tcp", ":8888")
		if err != nil {
			fmt.Println("listen failed")
		}
		rpcServer := grpc.NewServer()
		RegisterGreeterServer(rpcServer, &RPCService{})
		if err := rpcServer.Serve(lis); err != nil {
			fmt.Printf("rpc failed: %v", err)
		}
	})
	return nil
}

func StopRPCServer() error {
	if rpcServer == nil {
		log.Fatal("rpc not start")
	}
	rpcServer.Stop()
	return nil
}

func NewRPCServer() *grpc.Server {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen failed")
	}
	svr := grpc.NewServer()
	RegisterGreeterServer(svr, &RPCService{})
	if err := svr.Serve(lis); err != nil {
		fmt.Printf("rpc failed: %v", err)
	}

	return svr
}
