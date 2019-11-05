package main

import (
	"fmt"
	"time"
	// "github.com/EDDYCJY/gsema"
)

func main() {
	// timeStr := time.Now().Format("2006-01-02 15:04:56")
	t, _ := time.ParseInLocation("2006-01-02", "1994-03-01", time.Local)
	timeUnix := t.Unix()
	fmt.Println(timeUnix)
}
