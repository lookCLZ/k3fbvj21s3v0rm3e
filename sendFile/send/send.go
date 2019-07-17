package main

import (
	"os"
	"fmt"
	"net"
	"io"
)

func main() {
	args:=os.Args
	filePath:=args[1]

	fileInfo,_:=os.Stat(filePath)
	fileName:=fileInfo.Name()

	// 建立连接
	conn,_:=net.Dial("tcp","127.0.0.1:8008")
	defer conn.Close()

	// 写入数据(文件名）
	_,_=conn.Write([]byte(fileName))
	
	buf:=make([]byte,16)

	// 读取返回数据
	n,_:=conn.Read(buf)
	if "ok"==string(buf[:n]) {
		sendFile(conn,filePath)
	}
}

func sendFile(conn net.Conn, filePath string) {
	f,_:=os.Open(filePath)
	defer f.Close()

	buf:=make([]byte,4096)
	for {
		n,err:=f.Read(buf)
		if err!=nil{
			if err==io.EOF {
				fmt.Println("文件发送完成")
			} else {
				fmt.Println("os.Open err:",err)
			}
			return
		}
		// 写入到接收端
		_,_=conn.Write(buf[:n])
	}
}
