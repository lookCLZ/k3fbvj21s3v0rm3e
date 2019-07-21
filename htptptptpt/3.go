package main

import (
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request) {
	// w:响应
	// r:获取
	w.Write([]byte("hello 9fud9fd9"))
}

func main() {
	// 注册回调函数，改回调函数汇总服务器被访问时，自动调用
	http.HandleFunc("/itcast",handler)
	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000",nil)
}