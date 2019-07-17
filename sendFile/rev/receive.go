package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	// 创建用于监听的socket
	listener,_:=net.Listen("tcp","127.0.0.1:8008")
	defer listener.Close()

	conn,_:=listener.Accept()
	defer conn.Close()

	buf:=make([]byte,4096)
	n,_:=conn.Read(buf)

	fileName:=string(buf[:n])
	conn.Write([]byte("ok"))

	recvFile(conn,fileName)
}

func recvFile(conn net.Conn,fileName string) {
	f,_:=os.Create(fileName)
	defer f.Close()

	buf:=make([]byte,4096)
	for{
		n,_:=conn.Read(buf)
		if n==0{
			fmt.Println("接收文件完成")
			return
		}
		f.Write(buf[0:n])
	}
}