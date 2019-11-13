package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	db *bolt.DB
	tail []byte 
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

func NewBlockChain() *BlockChain {
	var lastHash []byte 
	// 打开数据库
	db,err:=bolt.Open(blockChainDb,0600,nil)

	if err!=nil{
		log.Panic("打开数据库失败")
	}

	db.Update(func(tx *bolt.Tx)error {
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			bucket,err=tx.CreateBucket([]byte(blockBucket))
			if err!=nil{
				log.Panic("创建bucket失败")
			}

			genesisBlock:=GenesisBlock()

			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"),genesisBlock.Hash)
			lastHash = genesisBlock.Hash 

		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db,lastHash}
}

func GenesisBlock() *Block{
	return NewBlock("Go一期创世块，牛逼！",[]byte{})
}

func (bc *BlockChain) AddBlock(data string){
	db:=bc.db 
	lastHash:=bc.tail 

	db.Update(func(tx *bolt.Tx)error{
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			log.Panic("bucket 不应该为空，请检查")
		}

		block:=NewBlock(data,lastHash)
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte("LastHashKey"),block.Hash)

		bc.tail = block.Hash
		return nil
	})
}

