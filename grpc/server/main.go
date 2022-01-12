package main

import (
	hello_grpc "bitcoin/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
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
	for {
		recv, err := hi1Server.Recv()
		fmt.Println(recv)
		if err != nil {
			hi1Server.SendAndClose(&hello_grpc.Res{Message: "OK", Age: 200})
			break
		}
	}
	return nil
}
func (s *server) SayHi2(req *hello_grpc.Req, hi2Server hello_grpc.HelloGRPC_SayHi2Server) error {
	message := req.Message
	age := req.Age
	fmt.Println(message)
	fmt.Println(age)
	i := 0
	for true {
		if i > 10 {
			break
		}
		time.Sleep(1 * time.Second)
		hi2Server.Send(&hello_grpc.Res{Message: message, Age: age})
		i++
	}
	return nil
}
func (s *server) SayHi3(hi3Server hello_grpc.HelloGRPC_SayHi3Server) error {
	mes := make(chan string)
	ag := make(chan int32)
	i := 0
	go func() {
		for true {
			recv, err := hi3Server.Recv()
			if err != nil {
				mes <- "end"
			}
			mes <- recv.Message
			ag <- recv.Age
		}
	}()
	for true {
		if i < 10 {
			s1 := <-mes
			s2 := <-ag
			if s1 == "end" {
				break
			}
			hi3Server.Send(&hello_grpc.Res{Message: s1, Age: s2})
			i++
		}
	}
	return nil
}

func main() {
	l, _ := net.Listen("tcp", ":9999")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
