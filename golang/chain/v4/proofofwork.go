package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

type ProofOfWork struct {
	block *Block 
	target *big.Int 
}

func NewProofOfWork(block *Block) *ProofOfWork{
	pow:=ProofOfWork{
		block:block,
	}
	targetStr:="0000100000000000000000000000000000000000000000000000000000000000"
	tmpInt:=big.Int{}
	tmpInt.SetString(targetStr,16)

	pow.target = &tmpInt
	return &pow
}

func (pow *ProofOfWork) Run() ([]byte,uint64) {
	var nonce uint64 
	block := pow.block 
	var hash [32]byte 
	fmt.Println("开始挖矿...")
	for {
		tmp:=[][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			// 只对区块头做哈希值，区块体通过MerkelRoot产生影响
			// block.Data
		}
		// 将二维的切片连接，返回一个一位切片
		blockInfo:=bytes.Join(tmp,[]byte{})
		hash=sha256.Sum256(blockInfo)
		tmpInt:=big.Int{}
		tmpInt.SetBytes(hash[:])

		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功！hash:%x,nonce:%d\n",hash,nonce)
			return hash[:],nonce
		} else {
			nonce++
		}
	}
}