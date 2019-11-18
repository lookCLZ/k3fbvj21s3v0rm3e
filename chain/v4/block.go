package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

//  定义结构
type Block struct {
	Version      uint64
	PrevHash     []byte
	MerkelRoot   []byte
	TimeStamp    uint64
	Difficulty   uint64
	Nonce        uint64
	Hash         []byte
	Transactions []*Transaction
}

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// 新建block
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	block := Block{
		Version:      00,
		PrevHash:     prevBlockHash,
		MerkelRoot:   []byte{},
		TimeStamp:    uint64(time.Now().Unix()),
		Difficulty:   0,
		Nonce:        0,
		Hash:         []byte{},
		Transactions: txs,
	}

	block.MerkelRoot = block.MakeMerkelRoot()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return &block
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码出错！")
	}
	return buffer.Bytes()
}

func Deserialize(data []byte) Block {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错！", err)
	}
	return block
}

func (block *Block) MakeMerkelRoot() []byte {
	return []byte{}
}
