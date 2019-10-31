package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"
)

// 定义结构
type Block struct {
	Version    uint64
	PrevHash   []byte
	MerkelRoot []byte
	TimeStamp  uint64
	Difficulty uint64
	Nonce      uint64

	Hash []byte
	Data []byte
}

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer

	/**
	 *binary.Write(w io.Writer, order binary.ByteOrder, data interface{})
	 *write将数据的二进制写入w
	 *数据必须是固定大小的值或者固定大小的片段值，或者指向此类数据的指针
	 *布尔值编码为一个字节：1表示真，0表示假
	 *写入w的字节使用指定的字节顺序进行编码
	 *并从数据的连续字段中读取
	 *在编写结构时，字段的值为零
	 **/
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	/**
	*返回一个长度为b.len()的片段，其中包含缓冲区的未读部分
	*slice只在下一次缓冲区修改，即直到下一次调用read、write、reset或者truncate方法
	*slice至少在下一次修改缓冲区之前给缓冲区内容取别名，
	*因此切片的立即更改将影响将来读取的结果
	*********/
	return buffer.Bytes()
}

// 创建区块
func NewBlock(data string,prevBlockHash []byte) *Block{
	block:=Block{
		Version: 00,
		PrevHash: prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce: 0,
		Hash: []byte{},
		Data: []byte(data),
	}

	// 创建pow对象
	pow:=NewProofOfWork(&block)
	// 查找随机数，不停的进行hash运算
	hash,nonce:=pow.Run()

	// 根据挖矿的结果对区块数据进行更新
	block.Hash = hash 
	block.Nonce = nonce 

	return &block
}