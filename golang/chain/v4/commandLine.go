package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	fmt.Printf("添加区块链成功!\n")
}

func (cli *CLI) AddBlock(data string) {
	fmt.Printf("添加区块链成功！\n")
}

func (cli *CLI) PrinBlockChainReverse() {
	bc := cli.bc
	// 创建迭代器
	it := bc.NewIterator()
	// 调用迭代器
	for {
		block := it.Next()
		fmt.Printf("================\n\n")
		fmt.Printf("版本号：%d\n", block.Version)
		fmt.Printf("当前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("梅克尔根：%x\n", block.MerkelRoot)
		fmt.Printf("时间戳：%d\n", block.TimeStamp)
		fmt.Printf("难度值：%d\n", block.Difficulty)
		fmt.Printf("随机数：%d\n", block.Nonce)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Transactions[0].TXInputs[0].Sig)

		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束")
			break
		}
	}
}
