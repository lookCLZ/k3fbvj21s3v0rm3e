package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

// 服务器用于实现helloworld.GreeterServer。
type server struct{}

// SayHello实现helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello Everyone" + in.Name}, nil
}
func main() { //监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//new服务对象
	s := grpc.NewServer()
	//注册服务
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务器上注册反射服务。 reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
