package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

func paddingLastGroup(plainText []byte, blockSize int) []byte {
	// 1.获取最后一个组中剩余的字节数 (1-8)
	padNum := blockSize - len(plainText)%blockSize
	// 2.创建新的切片，长度 == padNum, 每个字节值byte(padNum)
	char := []byte{byte(padNum)} // 长度1
	newPlain := bytes.Repeat(char, padNum)
	// 3.newPlain数组追加到原始明文的后边
	newText := append(plainText, newPlain...)

	return newText
}

func unPaddingLastGroup(plainText []byte) []byte {
	// 获取去切片中最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	// 获取尾部填充的字节个数
	number := int(lastChar)

	return plainText[:length-number]
}

// des加密 使用CBC分组模式
func desEncrypt(plainText, key []byte) []byte {
	// 建立一个des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 明文填充
	newText := paddingLastGroup(plainText, block.BlockSize())
	// 创建一个CBC分组模式加密的接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 加密
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)

	return cipherText
}

// des解密 使用CBC分组模式
func desDecrypt(cipherText, key []byte) []byte {
	// 建立一个des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 创建一个CBC分组模式解密的接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 解密
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(cipherText, cipherText)
	// 删除加密时填充的尾部数据
	plainText = unPaddingLastGroup(cipherText)

	return plainText
}

// aes加密 使用CTR分组模式
func aesEncrypt(plainText, key []byte) []byte {
	// 建立一个aes的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 创建一个CTR分组模式加密的接口
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	// 加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

// aes解密 使用CTR分组模式
func aesDecrypt(cipherText,key []byte)[]byte{
	// 建立一个aes的密码接口
	block,err:=aes.NewCipher(key)
	if err!=nil{
		panic(err)
	}
	// 创建一个使用CTR分组模式解密的接口
	iv:=[]byte("12345678abcdefgh")
	stream:=cipher.NewCTR(block,iv)
	// 解密
	plainText:=make([]byte,len(cipherText))
	stream.XORKeyStream(plainText,cipherText)

	return plainText
}
