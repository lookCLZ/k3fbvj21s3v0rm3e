package main

type BlockChain struct {
	// 定一个区块链数组
	blocks []*Block
}

// 区块链构造函数
func NewBlockChain() *BlockChain {
	// 创建一个创世块
	genesisBlock:=GenesisBlock()
	return &BlockChain{
		blocks:[]*Block{genesisBlock},
	}
}

// 创世块函数
func GenesisBlock() *Block{
	return NewBlock("GO一期创世块，老牛逼了!",[]byte{})
}

// 添加区块
func (bc *BlockChain) AddBlock(data string) {
	// 先获取最后一个区块，然后得到其哈希值，作为新增一个区块的前置哈希
	lastBlock:=bc.blocks[len(bc.blocks)-1]
	preHash:=lastBlock.Hash

	// 创建新的区块
	block:=NewBlock(data,preHash)
	// 将新的区块拼接到末尾
	bc.blocks = append(bc.blocks,block)
}
