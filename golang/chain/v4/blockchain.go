package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	db   *bolt.DB
	tail []byte
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

func NewBlockChain(address string) *BlockChain {
	var lastHash []byte
	db, err := bolt.Open(blockChainDb, 0600, nil)

	if err != nil {
		log.Panic("打开数据库失败！")
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}

			genesisBlock := GenesisBlock(address)
			fmt.Printf("genesisBlock:%s\n", genesisBlock)

			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

func (bc *BlockCHain) PrintChain() {
	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockBucket"))

		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}

			block := Deserialize(v)

			fmt.Printf("=========区块高度：%d==========\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号：%d\n", block.Version)
			fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
			fmt.Printf("梅克尔根：%x\n", block.MerkelRoot)
			fmt.Printf("时间戳：%d\n", block.TimeStamp)
			fmt.Printf("难度值：%d\n", block.Difficulty)
			fmt.Printf("随机数：%d\n", block.Nonce)
			fmt.Printf("当前区块哈希值：%x\n", block.Hash)
			fmt.Printf("区块数据：%d\n", block.Transactions[0].TXInputs[0].Sig)
			fmt.Printf("时间戳：%d\n", block.TimeStamp)
			return nil
		})
		return nil
	})
}

func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	// 定义一个map来保存消费过的output，key是这个output的交易id,value是这个交易中索引的数组
	// map[交易id][]int64
	spendOutputs := make(map[string][]int64)

	it := bc.NewIterator()

	for {
		block := it.Next()
		for _, tx := range block.Transactions {
			fmt.Printf("current txid: %x\n", tx.TXID)

		OUTPUT:
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index: %d\n", i)

				// 将所有消耗过的output和当前的所即将添加output对比一下
				if spendOutputs[string(tx.TXID)] != nil {
					for _, j := range spendOutputs[string(tx.TXID)] {
						if int64(i) == j {
							continue OUTPUT
						}
					}
				}

				// 这个output和我们目标的地址相同，满足条件，添加到返回的UTXO数组中
				if output.PubKeyHash == address {
					UTXO = append(UTXO, output)
				}
			}

			// 如果是挖矿交易，不做遍历直接跳过
			if !tx.IsCoinbase() {
				// 遍历input,找到自己花费过的utxo的集合
				for _, input := range tx.TXInputs {
					// 判断当前这个input和目标（李四）是否一致，如果相同，说明是李四消耗过的output，加进来
					if input.Sig == address {
						spendOutputs[string(input.TXid)] = append(spendOutputs[string(input.TXid)], input.Index)
					}
				}
			} else {
				fmt.Printf("这个是coinbase，不做input遍历！")
			}
		}

		if len(block.PrevHash) == 0 {
			break
			fmt.Printf("区块比那里完成退出！")
		}
	}
	return UTXO
}

func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]uint64, float64) {
	var utxos map[string][]uint64
	var calc float64

	return utxos, calc
}
