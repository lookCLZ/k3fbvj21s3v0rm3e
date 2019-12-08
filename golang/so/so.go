package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

var cond sync.Cond   // 定义条件变量

func producer(out chan<- int,idx int) {
	for {
		cond.L.Lock()
		// 判断缓冲区是否满
		if len(out)==5{
			cond.Wait()   //1.阻塞，2.释放锁, 3.唤醒后重新加锁
		}
		num:=rand.Intn(800)
		out<-num
		fmt.Printf("生产者%d,生产:%d\n",idx,num)
		// 访问公共区结束，且打印结束后，解锁
		cond.L.Unlock()
		// 唤醒阻塞在条件变量上的对端
		cond.Signal()
		time.Sleep(time.Millisecond*200)
	}
}

func customer(in <-chan int,idx int) {
	for {
		cond.L.Lock()
		if len(in)==0{
			cond.Wait()
		}
		num:=<-in
		fmt.Printf("---消费者%d,消费者:%d\n",idx,num)
		cond.L.Unlock()
		cond.Signal()
	}
}

func main() {
	product:=make(chan int)	    //创建一个channel
	rand.Seed(time.Now().UnixNano())

	// 指定条件变量使用的锁
	cond.L=new(sync.Mutex)  //初始值为0，未加锁状态
	
	// 创建5个Go程，也就是5个生产者
	for i:=0;i<5;i++{
		go producer(product,i+1)
	}

	// 创建5个Go程，也就是5个消费者
	for i:=0;i<5;i++{
		go customer(product,i+1)
	}



	for ;;{

	}
}