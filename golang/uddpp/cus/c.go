package main

import (
	"fmt"
	"net"
)

func  main()  {
	// 创建套接字
	conn,_:=net.Dial("udp","127.0.0.1:8006")
	_,_=conn.Write([]byte("你吃草区别"))
	defer conn.Close()

	buf:=make([]byte,4096)
	for i:=0;i<10000000;i++{
		conn.Write([]byte("你吃草去吧"))
		n,err:=conn.Read(buf)
		if n==0 {
			fmt.Println("退出了啊")
			return
		}
		if err!=nil{
			fmt.Println("发生错误")
			return
		}
		fmt.Println("服务器返回：",string(buf[:n]))
	}
}