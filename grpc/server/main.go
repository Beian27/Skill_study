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
	return &hello_grpc.Res{
		Message: "服务端返回的grpc的内容",
	}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":9999")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
