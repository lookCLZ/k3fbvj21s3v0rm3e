package main

import (
	"fmt"
	"time"
	"runtime"
)

func aaa()  {
	for {
		time.Sleep(time.Millisecond * 200)
		fmt.Println("-----------")
	}
}

func main()  {
	go func() {
		fmt.Println("------------1")
		go aaa()
		fmt.Println("------------2")
		return
	}()

	for {
		// GC执行一次垃圾回收。并且阻塞，直到垃圾回收完成
		// 会等待循环终止
		runtime.GC();
	}
}