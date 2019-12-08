package main

import (
	"net/http"
	"fmt"
	// "io"
)

func main() {
	resp,_:=http.Get("http://www.baidu.com")
	defer resp.Body.Close()

	buf:=make([]byte,4096)
	var result string

	for {
		n,_:=resp.Body.Read(buf)	
		if n==0{
			fmt.Println("读取完毕")
			break	//这里不能用return，否则直接就退出main函数了
		}
		// fmt.Println(result)
		result+=string(buf[:n])
		// fmt.Printf("%s", string(buf[:n]))
	}
	fmt.Println("result = ",result)
}