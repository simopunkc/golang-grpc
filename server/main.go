package main

import (
	"context"

	"golang-grpc/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func (s *server) Add(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}
