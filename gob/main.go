package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	// "github.com/astaxie/beego/logs"
	// "github.com/betterjun/pkg/logs"
	log "github.com/sirupsen/logrus"
)

type Person struct {
	Name string
	Age  uint
}

var logN = log.New()

func main() {
	var xiaoming, daming Person
	xiaoming.Name = "晓明"
	xiaoming.Age = 20

	var buffer bytes.Buffer

	// 创建一个编码器
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(&xiaoming)
	log.Printf("编码后的晓明：%v\n", buffer.Bytes())

	// 创建一个解码器
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	_ = decoder.Decode(&daming)
	fmt.Println(daming)
}
