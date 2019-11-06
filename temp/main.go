package main

import (
	"fmt"
	"strings"
	"time"
	// "github.com/EDDYCJY/gsema"
)

func main() {
	// timeStr := time.Now().Format("2006-01-02 15:04:56")
	t, _ := time.ParseInLocation("2006-01-02", "1994-03-01", time.Local)
	timeUnix := t.Unix()
	fmt.Println(timeUnix)

	a := fmt.Sprint([]int{1, 2, 3})
	fmt.Println(a)
	a=strings.Trim(a, "[]")
	fmt.Println(a)
	a=strings.Replace(a," ",",",-1)
	fmt.Println(a)
	var temp []int 
	temp=append(temp,3)
	fmt.Println(temp)
}
