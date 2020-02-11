package main

import (
	"google.golang.org/grpc"
	"fmt"
	pd "github.com/grpc/day2/myproto"
	"context"
)

func main() {
	// create GRPC connection
	conn,err:=grpc.Dial("127.0.0.1:10086",grpc.WithInsecure())
	if err!=nil{
		fmt.Println("网络异常",err)
	}
	// network delay shutdown
	defer conn.Close()

	// create GPRC handle
	c:=pd.NewHelloserverClient(conn)
	// 通过句柄调用函数
	re,err:=c.SayHello(context.Background(),&pd.HelloReq{Name:"熊猫"})
	if err!=nil{
		fmt.Println("调用sayhello失败")
	}
	fmt.Println("调用sayhello的返回",re.Msg)

	re1,err:=c.Sayname(context.Background(),&pd.NameReq{Name:"托尼斯塔克"})
	if err!=nil{
		fmt.Println("say name 调用失败")
	}
	fmt.Println("调用Sayname的返回",re1.Msg)


}