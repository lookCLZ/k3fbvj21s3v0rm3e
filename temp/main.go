package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	var content = "你好"

	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(content)

	data := buffer.Bytes()
	fmt.Println(string(data))
}
