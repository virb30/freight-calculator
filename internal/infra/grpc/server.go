package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	ServerPort string
	GrpcServer *grpc.Server
}

func NewGrpcServer(serverPort string) *GrpcServer {
	grpcServer := grpc.NewServer()
	server := &GrpcServer{
		ServerPort: serverPort,
		GrpcServer: grpcServer,
	}
	return server
}

func (s *GrpcServer) Start(ch chan error) {
	reflection.Register(s.GrpcServer)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.ServerPort))
	if err != nil {
		ch <- err
	}
	ch <- s.GrpcServer.Serve(lis)
}
