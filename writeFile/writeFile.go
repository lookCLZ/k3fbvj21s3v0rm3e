package main 

import (
	"fmt"
	"os"
	"io"
)

func main() {
	file1,err:=os.Open("./file.txt")
	if err!=nil{
		fmt.Println(err)
	}
	file2,err:=os.Create("./file2.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer func () {
		file1.Close()
		file2.Close()
		fmt.Println("关闭文件")
	}()

	buffer:=make([]byte,2)
	for {
	n,err:=file1.Read(buffer)
		if err==io.EOF {
			break
		}
	fmt.Println((buffer))
	file2.WriteString(string(buffer)[:n])
	}
}