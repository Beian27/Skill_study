package main

import (
	hello_grpc "bitcoin/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	fmt.Println(err)
	defer conn.Close()
	client := hello_grpc.NewHelloGRPCClient(conn)
	//res, err := client.SayHi(context.Background(), &hello_grpc.Req{
	//	Message: "张宏宇",
	//	Age:     22,
	//})

	//res, err := client.SayHi1(context.Background())
	//i := 0
	//for true {
	//	if i > 10 {
	//		err := res.CloseSend()
	//		fmt.Println(err)
	//		break
	//	}
	//	time.Sleep(1 * time.Second)
	//	res.Send(&hello_grpc.Req{Message: "age", Age: int32(i)})
	//	i++
	//}

	//res, err := client.SayHi2(context.Background(), &hello_grpc.Req{Message: "ZHY", Age: 200})
	//if err != nil {
	//	return
	//}
	//for true {
	//	recv, err := res.Recv()
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println(recv)
	//}

	res, err := client.SayHi3(context.Background())
	if err != nil {
		return
	}
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		for true {
			err2 := res.Send(&hello_grpc.Req{Message: "ZHY", Age: 200})
			if err2 != nil {
				waitGroup.Done()
				break
			}
		}
	}()

	go func() {
		for true {
			time.Sleep(time.Second * 1)
			recv, err := res.Recv()
			if err != nil {
				waitGroup.Done()
				break
			}
			fmt.Println(recv)
		}
	}()
	waitGroup.Wait()

	fmt.Println(res)
}
