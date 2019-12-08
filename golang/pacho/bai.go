package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var count = make(chan int)
func HttpGet(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	var result string

	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break //这里不能用return，否则直接就退出main函数了
		}
		result += string(buf[:n])
	}
	return result
}

func working(start, end int) {
	fmt.Printf("正在爬取第%d页到第%d页...\n", start, end)
	// 循环爬取每一页数据
	for i := start; i <= end; i++ {
		go func(i int) {
			url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
			res := HttpGet(url)

			f, _ := os.Create("第" + strconv.Itoa(i) + "页.html")
			f.WriteString(res)
			fmt.Println("第",i,"页抓取完成")
			defer f.Close()
			count <- i
		}(i)
	}
}

func main() {
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入起始页")
	input, _ := inputReader.ReadSlice('\n')
	temp := input[:len(input)-1]
	inputNumStart, _ := strconv.Atoi(string(temp))

	fmt.Println("请输入结束页")
	input, _ = inputReader.ReadSlice('\n')
	temp = input[:len(input)-1]
	inputNumEnd, _ := strconv.Atoi(string(temp))

	working(inputNumStart, inputNumEnd)

	for i:=inputNumStart;i<=inputNumEnd;i++ {
		<-count
	}
}


