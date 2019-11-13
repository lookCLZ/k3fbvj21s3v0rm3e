package main


import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

// 定义一个工作量证明的接口ProofOfWork
type ProofOfWork struct {
	block *Block 
	target *big.Int
}

// 构造函数
func NewProofOfWork(block *Block)*ProofOfWork{
	pow:=ProofOfWork{
		block:block,
	}
	// 指定难度值
	targetStr:="0000100000000000000000000000000000000000000000000000000000000000"
	var tmpInt big.Int
	tmpInt.SetString(targetStr,16)
	pow.target = &tmpInt
	return &pow
}

func (pow *ProofOfWork) Run() ([]byte,uint64){
	var nonce uint64 
	block:=pow.block 
	var hash [32]byte 

	fmt.Println("开始挖矿。。。")
	for {
		tmp:=[][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}

		blockInfo:=bytes.Join(tmp,[]byte{})

		hash = sha256.Sum256(blockInfo)
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