package main

import (
	"fmt"
	"os"
)

// 这是一个用来接收命令行参数并且控制区块链
type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA	"add data to blockchain"
	printChain				"print all blockchain data"
`

// 接受参数的动作，放到一个函数中
func (cli *CLI) Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}

	// 分析命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加区块\n")
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Println("添加区块链参数使用不当，请检查")
			fmt.Print(Usage)
		}
	case "printChain":
		fmt.Println("打印区块")
		cli.PrinBlockChain()
	default:
		fmt.Printf("无效的命令，请检查")
		fmt.Printf(Usage)
	}
}
