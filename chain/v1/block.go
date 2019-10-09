package main

import (
	"crypto/sha256"
	"time"
	"bytes"
	"encoding/binary"
	"log"
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
	Nonce	uint64

	// a.当前区块hash，正常比特币区块中没有当前区块的hash
	Hash	[]byte 
	// b.数据
	Data	[]byte
}

// 1.补充区块字段
// 2.更新计算hash函数
// 3.优化代码

// 实现一个分支函数，功能是将uint64转换成[]byte
func Uint64ToByte(num uint64)[] byte{
	var buffer bytes.Buffer

	err:=binary.Write(&buffer,binary.BigEndian,num)
	if err!=nil{
		log.Panic(err)
	}
	return buffer.Bytes()
}