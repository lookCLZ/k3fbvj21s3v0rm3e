package main

import (
	"fmt"

	// myproto "./text"
	"github.com/golang/protobuf/proto"
)

func main() {
	test := &Test{
		Name:    "熊猫",
		Stature: 180,
		Weight:  []int64{120, 125, 198, 180, 150, 180},
		Motto:   "一盘臭气",
	}
	// 将Struct转换成protobuf
	data, _ := proto.Marshal(test)
	// fmt.Println(data)
	// 创建一个新的Test结构体
	newTest := &Test{}
	// 将data转换为test结构体
	_ = proto.Unmarshal(data, newTest)
	fmt.Println(newTest.String())
	fmt.Println(newTest.GetName())
	fmt.Println(newTest.GetWeight())
}
