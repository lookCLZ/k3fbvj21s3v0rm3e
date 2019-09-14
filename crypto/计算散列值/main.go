package main

import (
	// "bytes"
	"fmt"
	"crypto/md5"
	"crypto/sha256"
)

func main() {
	str:="天龙八部"
	hash:=md5.New()
	hash.Write([]byte(str))
	hash.Write([]byte(str))

	result:=hash.Sum(nil)
	res:=fmt.Sprintf("%x",result)
	fmt.Println(res)
	result2:=sha256.Sum256([]byte(str))
	res2:=fmt.Sprintf("%x",result2)
	fmt.Println(res2)
}
