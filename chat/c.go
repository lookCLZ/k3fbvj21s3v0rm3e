package main

import (
	"net"
	"fmt"
	// "strings"
	// "time"
)

// 客户端信息
type Client struct {
	C chan string
	Name string 
	Addr string
}

// 在线统计
var onlineMap = make(map[string]Client)

// 消息
var message = make(chan string)

func main() {
	listener,_:=net.Listen("tcp","127.0.0.1:8005")
	defer listener.Close()

	go Manager()
	for {
		conn,_:=listener.Accept()
		go HandlerConnect(conn)
	}
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	netAddr:=conn.RemoteAddr().String()

	clnt:=Client{
		C:make(chan string),
		Name:netAddr,
		Addr:netAddr,
	}
	onlineMap[netAddr]=clnt

	// go WriteMsgToClient(conn,clnt)
	fmt.Println(1)
	message <- MakeMsg(clnt,"已上线")
	fmt.Println(2)
	go ReadMsg(conn,clnt)
}

// 全go程管理器
func Manager() {
	for {
		msg:=<-message
		for _,v:=range onlineMap {
			v.C<-msg
		}
	}
}

// 监听自己的耳朵，有就写入
func WriteMsgToClient(conn net.Conn,clnt Client) {
	for msg:=range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

// 组装用户的消息的格式
func MakeMsg(clnt Client,msg string) (string) {
	return clnt.Name + msg
}

func ReadMsg(conn net.Conn,clnt Client) {
	buf:=make([]byte,4096)
	for {
		n,_:=conn.Read(buf)
		if n==0{ 
			message <- MakeMsg(clnt, "刚刚下线了")
			return
		}

		msg:=string(buf)
		message <- MakeMsg(clnt, msg)

	}
}

