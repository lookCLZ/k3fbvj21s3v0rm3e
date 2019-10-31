package main

import "fmt"

func main() {
	// 创建一个区块链
	bc := NewBlockChain()
	// 添加一个区块
	bc.AddBlock("班长向班花转了50枚比特币")
	// 再添加一个区块
	bc.AddBlock("班长又向班花转了50枚比特币")
	// 遍历这个区块链
	for i, block := range bc.blocks {
		fmt.Println("=======当前区块高度：%d========\n", i)
		fmt.Println("前一个区块哈希值：%\n", block.PrevHash)
		fmt.Println("当前区块哈希值：%s\n", block.Data)
		fmt.Println("区块数据：%s\n", block.Data)
	}
}
