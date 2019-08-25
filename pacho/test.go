package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, e := http.Get("https://movie.douban.com/top250?start=0&filter=")
	fmt.Println(e)
	fmt.Println(r)
	buf := make([]byte, 4096)
	var res string
	for {
		n, _ := r.Body.Read(buf)
		if n == 0 {
			break
		}
		res += string(buf[:n])
	}
	fmt.Println(res)
}
