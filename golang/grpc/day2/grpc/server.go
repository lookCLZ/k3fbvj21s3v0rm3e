package main

import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/grpc/day2/myproto"

	"context"
)

type server struct {}

func (this *server)SayHello(ctx context.Context,in *myproto.HelloReq) (out *myproto.HelloRsp,err error) {
	return &myproto.HelloRsp{Msg:"hello"+in.Name},nil
}

func (this *server)Sayname(ctx context.Context,in *myproto.NameReq) (out *myproto.NameRsp,err error) {
	return &myproto.NameRsp{Msg:in.Name+"早上好"},nil
}

func main() {
	// create monitoring
	ln,err:=net.Listen("tcp",":10086")
	if err!=nil{
		fmt.Println("网络错误",err)
	}

	// create GRPC server
	srv:=grpc.NewServer()

	// register server
	myproto.RegisterHelloserverServer(srv,&server{})

	// waiting connection
	err=srv.Serve(ln)
	if err!=nil{
		fmt.Println("网络错误",err)
	}
}
