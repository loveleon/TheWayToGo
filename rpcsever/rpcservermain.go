package main

import (
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"golang.org/x/net/context"
	"fmt"
	"net"
	"google.golang.org/grpc"
	//"github.com/grpc/grpc-go/examples/helloworld/helloworld"
	//"reflect"
	"google.golang.org/grpc/reflection"
	//"github.com/grpc/grpc-go"
	"time"
)

const (
	port = "127.0.0.1:50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest )(*pb.HelloReply,error){
	fmt.Println("call server SayHello.")
	return &pb.HelloReply{"Sever'SayHello is called." + request.Name},nil
}

func (s *server) GetServerTime(ctx context.Context, request *pb.TimeRequest)(*pb.TimeReply,error){
	t := time.Now()
	//timereply := int64(t.Nanosecond())
	timereply := int64(t.Year())
	fmt.Println("GetSererTime is called.")

	return &pb.TimeReply{timereply},nil
}

func  main()  {
	//rpc server create .
	//Listen at port
	lis,err := net.Listen("tcp",port)
	if err != nil {
		fmt.Println("listen at ",port,"error.")
	}

	//create server
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println("create grpc server error.")
	}
}
