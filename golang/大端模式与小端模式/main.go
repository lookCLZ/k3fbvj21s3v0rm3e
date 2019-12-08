package main

import (
	"fmt"
	"time"
	"unsafe" //go语言的sizeof
)

func main() {
	s := int16(0x1234)
	b := int8(s)
	//0x1234
	// 低 --------》 高
	// 12 34  -> 大端 -> 高尾端
	// 34 12  -> 小端 -> 低尾端

	fmt.Println("int16字节大小为", unsafe.Sizeof(s)) //结果为2
	fmt.Println()
	if 0x34 == b {
		fmt.Println("little endian")
	} else {
		fmt.Println("big endian")
	}

	// 时间戳转日期格式
	t := time.Now().Unix()
	fmt.Println("t:", t)
	t2 := time.Unix(t, 0).Format("2006-01-02 15:04:05")
	fmt.Println("t2:", t2)
	// 日期格式转时间戳
	t3, _ := time.Parse("2006-01-02 15:04:05", t2)
	fmt.Println("t3:", t3)
	t4 := t3.Unix()
	fmt.Println("t4:", t4)
	// timeFormat := time.Unix(int64(time.Unix().Second()), 0).Format("2006-01-02 15:04:05")
	// fmt.Printf("时间戳: %s\n", timeFormat)
}
