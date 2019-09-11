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
	RPCServ *RPCService
	RPCOnce   sync.Once
)

type RPCService struct {
	StopCh chan struct{}
	GrpcSv *grpc.Server
}

func (s *RPCService) StartJob(ctx context.Context, in *JobRequest) (*JobReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &JobReply{Message: "Hello " + in.GetName()}, nil
}

func (s *RPCService) StopJob(ctx context.Context, in *JobRequest) (*JobReply, error) {
	log.Printf("Received Agagin: %v", in.GetName)
	close(s.StopCh)
	return &JobReply{Message: "Hello Agagin " + in.GetName()}, nil
}

func StartRPCServer() error {
	fmt.Println("RPC SERVER START")

	RPCServ = &RPCService{
		StopCh: make(chan struct{}),
		GrpcSv: grpc.NewServer(),
	}
	//rpcOnce.Do(func() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen failed")
	}
	RegisterGreeterServer(RPCServ.GrpcSv, RPCServ)
	if err := RPCServ.GrpcSv.Serve(lis); err != nil {
		fmt.Printf("rpc failed: %v", err)
	}
	/*
		rpcServer := grpc.NewServer()
		RegisterGreeterServer(rpcServer, &RPCService{stopCh: make(chan struct{})})
		if err := rpcServer.Serve(lis); err != nil {
			fmt.Printf("rpc failed: %v", err)
		}
	*/
	//})

	return nil
}

func StopRPCServer() error {
	if RPCServ == nil {
		log.Fatal("rpc not start")
	}
	RPCServ.GrpcSv.Stop()
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
