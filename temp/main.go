package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte{2,3}
	fmt.Println(a)
	a=bytes.Repeat(a, 6)
	fmt.Println(a)
}
