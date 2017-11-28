package main

import (
	//"net"
	"fmt"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc"
	"context"
	//"google.golang.org/grpc/reflection"
	"time"
)

const (
	remoteserver = "127.0.0.1:50051"
	defaultname = " defaultname"
)

func main(){
	//
	conn,err := grpc.Dial(remoteserver,grpc.WithInsecure())
	if err != nil {
		fmt.Println("connect to remote server ",remoteserver, "failed.")
		return
	}

	c := pb.NewGreeterClient(conn)

	reply,err := c.SayHello(context.Background(),&pb.HelloRequest{defaultname})
	if err != nil {
		fmt.Println("HelloReply get error.")
	}
	fmt.Println("hello reply: ",reply.GetMessage())
	time.Sleep(5*time.Second)

	timereply,err := c.GetServerTime(context.Background(),&pb.TimeRequest{defaultname})
	if err != nil {
		fmt.Println("time reply error.")
	}
	fmt.Println("time reply :",timereply.GetMessage())
	time.Sleep(2*time.Second)
	return
}
