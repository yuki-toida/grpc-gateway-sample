package main

import (
	"context"
	"net"

	pb "github.com/yuki-toida/grpc-gateway-sample/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, &EchoServiceServer{})
	server.Serve(listener)
}

type EchoServiceServer struct{}

func (s *EchoServiceServer) Post(c context.Context, m *pb.Message) (*pb.Message, error) {
	return m, nil
}

func (s *EchoServiceServer) Get(c context.Context, p *pb.Param) (*pb.Param, error) {
	return p, nil
}
