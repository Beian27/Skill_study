package main

import (
	hello_grpc "bitcoin/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	fmt.Println("1111111111111")
	fmt.Println(err)
	fmt.Println("1111111111111")
	defer conn.Close()
	client := hello_grpc.NewHelloGRPCClient(conn)
	res, err := client.SayHi(context.Background(), &hello_grpc.Req{Message: "客户端"})
	if err != nil {
		return
	}
	fmt.Println(res)
}
