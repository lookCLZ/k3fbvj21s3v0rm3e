package main

import (
	"net"
	"fmt"
	"os"
)

func errFunc(err error,info string) {
	if err!=nil {
		fmt.Println(info,err)
		// runtime.Goexit() //结束当前Go程
		os.Exit(1)  //结束当前进程
	}
}
// general
func main() {
	// 创建一个socket
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	errFunc(err,"aaa")
	defer listener.Close()

	// 阻塞等待客户端连接
	conn,_:=listener.Accept()
	defer conn.Close()

	buf:=make([]byte,4096)
	var n int
	
		n,err=conn.Read(buf)
		
		if n==0{
			fmt.Println(1)
			return
		}
		fmt.Println(err)
		fmt.Println(n)
		if err!=nil{
			fmt.Println(2)
			return
		}
	
	fmt.Println("//////////////////")
	fmt.Printf("|%s|\n",string(buf[:n]))

}
// 请求行
// 请求头
// 空行
// 请求包体
