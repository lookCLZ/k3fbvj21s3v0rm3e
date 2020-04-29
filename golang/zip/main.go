//main.go
package main

import (
    "flag"
    "fmt"
)

const help = `
程序名称: max
描述: 输出2个数,输出较大的数.
示例: 
1. test -help
2. max -first=5 -second=66
`
//
func main() {
    flag.Usage = func() {
        fmt.Print(help)
    }
    flag.Parse()
    //这里只实现使用帮助.
}