package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB 
	currentHashPointer []byte 
}

func (bc *BlockChain) NewIterator() *BlockChainIterator{
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

func (it *BlockChainIterator) Next() *Block{
	var block Block 
	it.db.View(func(tx *bolt.Tx)error{
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			log.Panic("迭代器遍历时bucket不应该为空")
		}
	})
}