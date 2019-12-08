package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {
	// data
	data := "hell"

	for i := 0; i < 1000; i++ {
		hash := sha256.Sum256([]byte(data + strconv.Itoa(i)))
		fmt.Println(strconv.Itoa(i))
		fmt.Printf("%x\n", hash)
	}
}
