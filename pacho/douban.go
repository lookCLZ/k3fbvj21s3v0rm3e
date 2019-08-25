package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// 指定爬取起始页，终止页
	var start, end int
	fmt.Println("请输入起始页")
	fmt.Scan(&start)
	fmt.Println("请输入终止页")
	fmt.Scan(&end)

	toWork(start, end)
	for {

	}
}

func toWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 页...\n", start, end)

	page := make(chan int) //防止主go程 提前结束

	for i := start; i < end+1; i++ {
		go SpiderPageDB(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Println("第%d页爬取完毕\n", <-page)
	}
}

func SpiderPageDB(idx int, page chan int) {
	var url = "https://movie.douban.com/top250?start=" + strconv.Itoa((idx-1)*25) + "&filter="

	result, _ := HttpGetDB(url)

	ret := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	filmName := ret.FindAllStringSubmatch(result, -1)

	ret = regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	filmScore := ret.FindAllStringSubmatch(result, -1)

	ret = regexp.MustCompile(`<span>(?s:(\d*?))人评价</span>`)
	filmPeople := ret.FindAllStringSubmatch(result, -1)

	Save2file(idx, filmName, filmScore, filmPeople)

	page <- idx
}

func HttpGetDB(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var buf = make([]byte, 4096)
	var n int
	// 循环爬取整页数据
	for {
		n, err = resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			return
		}
		result += string(buf[:n])
	}
	return
}

func Save2file(idx int, filmName, filmScore, filmPeople [][]string) {
	path := "/Users/liuhongrui/Downloads/未命名文件夹/" + strconv.Itoa(idx) + ".txt"
	f, _ := os.Create(path)
	defer f.Close()

	n := len(filmName)
	f.WriteString("电影名称" + "\t\t\t" + "评分" + "\t\t\t" + "人数" + "\n")
	for i := 0; i < n; i++ {
		f.WriteString(filmName[i][1] + "\t\t\t" + filmScore[i][1] + "\t\t\t" + filmPeople[i][1] + "\n")
	}
}
