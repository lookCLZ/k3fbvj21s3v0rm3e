package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	// 版本号
	Version uint64
	// 前区块hash
	PrevHash []byte
	// merkel根 （梅克尔根，这就是一个hash值，先不管）
	MerkelRoot []byte
	// 时间戳
	TimeStamp uint64
	// 难度值
	Difficulty uint64
	// 随机数,也就是挖矿要找的数据
	Nonce uint64

	// a.当前区块hash，正常比特币区块中没有当前区块的hash
	Hash []byte
	// b.数据
	Data []byte
}

// 1.补充区块字段
// 2.更新计算hash函数
// 3.优化代码

// 实现一个分支函数，功能是将uint64转换成[]byte
func Uint64ToByte(num uint64) []byte {
	// 定义一个byte buffer
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// Write writes the binary representation of data into w.
// Bytes returns a slice of length b.Len() holding the unread portion of the buffer.

// 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}

	block.SetHash()
	return &block
}

// 生成hash
func (block *Block) SetHash() {
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	// 将二维的切片数组连接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})

	// 2. sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
