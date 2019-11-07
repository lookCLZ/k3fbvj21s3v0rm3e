package main

import (
	"./blot"
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

	db.Update(func(tx *bolt.Tx)error) {
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			bucket,err=tx.CreateBucket([]byte(blockBucket))
			if err!=nil{
				log.Panic("创建")
			}
		}
	}
}

