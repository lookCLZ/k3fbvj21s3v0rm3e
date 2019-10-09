package main

import (
	"bytes"
	"log"
	"strings"
)

func main() {
	str := []string{"hello", "world", "!"}
	// 切片 =》 字符串
	res := strings.Join(str, "")
	log.Printf("res:%s\n", res)

	res1 := bytes.Join([][]byte{[]byte("hello"),[]byte("world")},[]byte(""))
	log.Printf("%s",res1)
}
