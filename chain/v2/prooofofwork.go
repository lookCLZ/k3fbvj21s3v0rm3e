package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 定义一个工作量证明的结构
type ProofOfWork struct {
	block *Block
	// 目标值
	// 一个非常大的数，有丰富的方法：比较，赋值
	target *big.Int
}

// 提供创建pow的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	// 我们指定的难度值，现在是一个string类型，需要进行类型转换
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)

	fmt.Println(tmpInt)
	pow.target = &tmpInt
	return &pow
}

// 提供计算不断计算hash的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var nonce uint64
	var block = pow.block
	var hash [32]byte

	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}

		// 将一个二维的数组连接起来，返回一个一维切片
		blockInfo := bytes.Join(tmp, []byte{})

		// 对这个二维数组做hash运算
		hash = sha256.Sum256(blockInfo)
		// 与pow中的target进行比较
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])

		// 比较hash
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功hash：%x ,nonce：%v\n", hash, nonce)
			return hash[:], nonce
		} else {
			// 没找到，继续挖
			nonce++
		}
	}
}
