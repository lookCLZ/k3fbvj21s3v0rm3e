// 4 x 3 = 12
package main

import (
	"io"
	"os"
	"fmt"
	"strings"
)

// count "hello" 
func RandAllTxt(src string,count *int) {
	file,err:=os.OpenFile(src,os.O_RDONLY,os.ModeDir)
	if err !=nil {
		fmt.Println("打开",src,"错误")
	}
	fileInfos,err:=file.Readdir(-1)
	if err!=nil {
		fmt.Println("读取目录",src,"错误")
	}
	for _,v:=range fileInfos {
		if v.IsDir() {
			RandAllTxt(src + "/" + v.Name(),count)
		} else if strings.HasSuffix(v.Name(),".txt") {
			file,err:=os.OpenFile(src + "/" + v.Name(),os.O_RDONLY,6)
			if err !=nil {
				fmt.Println("打开",src,"错误")
			}
			fmt.Println(file.Name())
			cap:=make([]byte,1024)
			for {
				fmt.Println(9)
				_,err:=file.Read(cap)
				if err==io.EOF {
					break
				}
				*count+=CountWord(cap,"hellow")
			}
		}
	}
}

func CountWord(textByte []byte,aim string) int {
	textStr:=string(textByte)
	strings.Fields(textStr)
	return strings.Count(textStr,aim)
}

func main() {
	var src string
	fmt.Println("请输入查找文件夹：")
	fmt.Scanf("%s",&src)
	fmt.Println("你输入的地址：",src)

	count:=0
	RandAllTxt(src,&count)
	fmt.Println("统计结果：",count)

}
