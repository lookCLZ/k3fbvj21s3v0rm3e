package main

import "fmt"

func main() {
	fmt.Println("des 加密解密 cbc分组模式")
	// 秘钥
	key:=[]byte("1234abdd")
	// 源文本
	src:=[]byte("你好，我是一段文本")
	cipherText:=desEncrypt(src,key)
	plainText:=desDecrypt(cipherText,key)
	fmt.Println("DES解密：",string(plainText))

	fmt.Println("aes 加密解密 ctr分组模式")
	// 秘钥
	key=[]byte("1234567812345678")
	// 源文本
	src=[]byte("世界，那么大")
	cipherText=aesEncrypt(src,key)
	plainText=aesDecrypt(cipherText,key)
	fmt.Println("AES解密：",string(plainText))
}