package main

import (
	"fmt"
	// "runtime"
)
func fib(ch chan int) {
	for {
		select {
		case num:=<-ch:
		fmt.Println(num)
		}
	}
}

func main() {
	ch:=make(chan int)
	x,y:=1,1

	go fib(ch)

	for i:=0;i<200;i++ {
		ch<-x
		x,y=y,y+x
	}
}