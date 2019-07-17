package main

import (
	"fmt"
	"net"
	"os"
	// "time"
)

func main() {
	// 指定服务器、IP Port 创建通信套接字
	conn,_:=net.Dial("tcp","127.0.0.1:8000")
	defer conn.Close()

	// 启用一个协程，从键盘中获取数据，将信息发送个服务器
	go func() {
		str:=make([]byte,4096)
		for {
			n,_:=os.Stdin.Read(str)
			conn.Write(str[:n])
		}
	}()

	buf:=make([]byte,4096)
	for {
		n,err:=conn.Read(buf)
		if n==0 {
			fmt.Println("监测到服务器关闭，客户端也将关闭")
			return
		}
		if err!=nil{
			fmt.Println("服务端端异常")
			return
		}
		fmt.Println("客户端读取到服务器回发:",string(buf[:n]))
	}
}