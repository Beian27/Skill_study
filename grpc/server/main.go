package main

import (
	hello_grpc "bitcoin/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {

	fmt.Println(req.GetMessage())
	fmt.Println(req.GetAge())
	return &hello_grpc.Res{
		Message: "Success",
		Age:     200,
	}, nil
}

func (s *server) SayHi1(hi1Server hello_grpc.HelloGRPC_SayHi1Server) error {
	return nil
}
func (s *server) SayHi2(req *hello_grpc.Req, hi2Server hello_grpc.HelloGRPC_SayHi2Server) error {
	return nil
}
func (s *server) SayHi3(hello_grpc.HelloGRPC_SayHi3Server) error {
	return nil
}

func main() {
	l, _ := net.Listen("tcp", ":9999")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
