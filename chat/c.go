package main

import (
	"fmt"
	"net"
	// "strings"
	"time"
)

// 客户端信息
type Client struct {
	C    chan string
	Name string
	Addr string
}

// 在线统计
var onlineMap = make(map[string]Client)

// 消息
var message = make(chan string)

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8005")
	defer listener.Close()

	go Manager()
	for {
		conn, _ := listener.Accept()
		go HandlerConnect(conn)
	}
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	// 判断客户端是否断开
	isQuit:=make(chan bool)
	run:=make(chan bool)

	netAddr := conn.RemoteAddr().String()

	clnt := Client{
		C:    make(chan string),
		Name: netAddr,
		Addr: netAddr,
	}
	onlineMap[netAddr] = clnt

	go WriteMsgToClient(conn, clnt)
	message <- MakeMsg(clnt, "已上线")
	fmt.Println(2)
	go ReadMsg(conn, clnt,isQuit,run)

	for {
		select {
		case <-isQuit:
			delete(onlineMap,clnt.Addr)
			return
		case <-run:
		case <-time.After(time.Second * 60):
			delete(onlineMap,clnt.Addr)
			return
		}
	}
}

// 全go程管理器
func Manager() {
	for {
		msg := <-message
		for _, v := range onlineMap {
			v.C <- msg
		}
	}
}

// 监听自己的耳朵，有就写入
func WriteMsgToClient(conn net.Conn, clnt Client) {
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

// 组装用户的消息的格式
func MakeMsg(clnt Client, msg string) string {
	return clnt.Name + "：" + msg
}

// 读取客户端发来的消息
func ReadMsg(conn net.Conn, clnt Client,isQuit,run chan<- bool) {
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			isQuit<-true
			message <- MakeMsg(clnt, "刚刚下线了")
			return
		}
		
		msg := string(buf[:n-1])
		run<-true
		if msg == "who" {
			fmt.Println(9)
			for _, v := range onlineMap {
				message <- "地址：" + v.Name
			}

		} else if len(msg) > 8 && msg[:6] == "rename" {
			clnt.Name = msg[7:]
			onlineMap[clnt.Addr] = clnt
			conn.Write([]byte("改名成功：" + clnt.Name + "\n"))
		} else {
			message <- MakeMsg(clnt, msg)
		}
	}
}
