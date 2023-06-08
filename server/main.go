package main

import (
	"context"
	"net"

	pb "github.com/josephchinedu/golang-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedAddServiceServer
}

func main() {

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	pb.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &pb.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &pb.Response{Result: result}, nil
}
