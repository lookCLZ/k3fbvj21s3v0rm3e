package main

import (
	"bytes"
	"fmt"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	db   *bolt.DB
	tail []byte
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

// 如果没有则创建，如果有则返回原来的
func NewBlockChain(address string) *BlockChain {
	var lastHash []byte
	db, _ := bolt.Open(blockChainDb, 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, _ = tx.CreateBucket([]byte(blockBucket))
			genesisBlock := GenesisBlock(address)
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

// 生成创世块
func GenesisBlock(address string) *Block {
	// 生成挖矿交易
	coinbase := NewCoinbaseTX(address, "Go一期创世块")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (bc *BlockChain) AddBlock(txs []*Transaction) {
	db := bc.db
	lastHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			fmt.Println("bucket不应该为空")
		}
		block := NewBlock(txs, lastHash)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)

		bc.tail = block.Hash
		return nil
	})
}

func (bc *BlockChain) Printchain() {
	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockBucket"))

		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}

			block := Deserialize(v)
			fmt.Println("============区块高度：%d============")
			blockHeight++
			fmt.Printf("版本号：%d\n", block.Version)
			fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
			fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
			fmt.Printf("时间戳: %d\n", block.TimeStamp)
			fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
			fmt.Printf("随机数 : %d\n", block.Nonce)
			fmt.Printf("当前区块哈希值: %x\n", block.Hash)
			fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Sig)
			return nil
		})
		return nil
	})
}

func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	txs := bc.FindUTXOTransactions(address)
	for _, tx := range txs {
		for _, output := range tx.TXOutputs {
			if address == output.PubKeyHash {
				UTXO = append(UTXO, output)
			}
		}
	}
	return UTXO
}

func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]uint64, float64) {
	utxos := make(map[string][]uint64)
	var calc float64
	txs := bc.FindUTXOTransactions(from)

	for _, tx := range txs {
		for i, output := range tx.TXOutputs {
			if from == output.PubKeyHash {
				if calc < amount {
					utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(i))
					calc += output.Value
					if calc >= amount {
						fmt.Println("找到了满足的金额：", calc)
						return utxos, calc
					}
				} else {
					fmt.Printf("不满足转账金额，当前总额：%f,目标金额：%f\n", calc, amount)
				}
			}
		}
	}
	return utxos, calc
}

func (bc *BlockChain) FindUTXOTransactions(address string) []*Transaction {
	var txs []*Transaction
	spentOutputs := make(map[string][]int64)

	it := bc.NewIterator()
	for {
		block := it.Next()
		for _, tx := range block.Transactions {
		OUTPUT:
			for i, output := range tx.TXOutputs {
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							continue OUTPUT
						}
					}
				}

				if output.PubKeyHash == address{
					txs=append(txs,tx)
				}

				if !tx.IsCoinbase() {
					for
				}
			}
		}
	}
}
