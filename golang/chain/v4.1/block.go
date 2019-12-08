package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

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

func Uint64ToByte(num int64) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

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
	encoder.Encode(&block)

	return buffer.Bytes()
}

func Deserialize(data []byte) Block {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block
	decoder.Decode(&block)

	return block
}

func (block *Block) MakeMerkelRoot() []byte {
	var info []byte
	for _, tx := range block.Transactions {
		info = append(info, tx.TXID...)
	}

	hash := sha256.Sum256(info)
	return hash[:]
}
