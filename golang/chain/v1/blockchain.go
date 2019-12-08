package main

// 引入区块链
type BlockChain struct {
	blocks []*Block
}

// 定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建第一个创世块，并作为第一个区块添加到区块链
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// 定义一个创世块
func GenesisBlock() *Block {
	return NewBlock("Go一期创世块，不错", []byte{})
}

// 添加区块
func (bc *BlockChain) AddBlock(data string) {
	// 获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	// 创建新的区块
	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks, block)
}
