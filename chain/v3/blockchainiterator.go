package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() {
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

func (it *BlockChainIterator) Next() {
	
}
