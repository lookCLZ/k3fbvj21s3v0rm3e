package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 指定服务器 通信协议，IP地址，port. 创建一个用于监听的socket
	listener,_:=net.Listen("tcp","127.0.0.1:8000")
	defer listener.Close()

	// 阻塞，等待客户端连接
	for{
		conn,_:=listener.Accept()

		fmt.Println("建立连接成功")
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	addr:=conn.RemoteAddr()
	buf:=make([]byte,4096)
	for {
		// 读取客户端发送过来的数据，可以从关闭的通道中读取数据
		n,_:=conn.Read(buf)
		if n==0 {
			fmt.Println("通道已经关闭球了")
			return
		}
		if "exit\n"==string(buf[:n]) {
			fmt.Println("客户端",addr,"挂逼了")
			return
		}
		str:=string(buf[:n])

		conn.Write([]byte(strings.ToUpper(str)))
		fmt.Println("服务器",addr,"读到数据",str)	
	}
}