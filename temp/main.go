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

	mmm:=testahan()
	fmt.Println(mmm)
	fmt.Println(*mmm)

	const (
		AutyTypeIdentity int = iota + 1 // 身份证
		AuthTypePassport int = iota              // 护照
		AuthTypePassport1 int = iota + 1
		AuthTypePassport2 int = iota
	)

	const bvb int = iota
	const bvb1 int = iota
	const bvb2 int = iota

	fmt.Println(AutyTypeIdentity)
	fmt.Println(AuthTypePassport)
	fmt.Println(AuthTypePassport1)
	fmt.Println(AuthTypePassport2)
	fmt.Println(bvb)
	fmt.Println(bvb1)
	fmt.Println(bvb2)
}

func testahan() *int {
	var a int 
	a = 2
	return &a
}