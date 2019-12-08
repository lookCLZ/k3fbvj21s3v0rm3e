package main

import (
	"fmt"
	"time"
	"net"
)

func main() {
	// 组织一个udp地址结构，指定服务器的IP+port
	serAddr,_:=net.ResolveUDPAddr("udp","127.0.0.1:8006")

	// 创建用于通信的socket
	udpConn,_:=net.ListenUDP("udp",serAddr)
	defer udpConn.Close()

	buf:=make([]byte, 4096)
	// 读取接收到的数据
	fmt.Print("0|\n")
	
	// 这里只有先读取客户端数据，获取获取客户端信息
	for {
		n,cltAddr,_:=udpConn.ReadFromUDP(buf)
	
		// 模拟处理数据
		fmt.Printf("服务器读到%v的数据：%s\n",cltAddr,string(buf[:n]))

		go func() {
			// 获取系统时间
			daytime:=time.Now().String()

			// 向客户端发送数据
			udpConn.WriteToUDP([]byte(daytime),cltAddr)		
		}()
	}
}
