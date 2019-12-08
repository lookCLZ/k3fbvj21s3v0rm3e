package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// 生成消息认证码
// HMAC是一种使用单向散列函数来构造消息认证码的方法
func GenerateHamc(plainText,key []byte) []byte{
	// 创建哈希接口，需要指定使用的哈希算法，和秘钥
	myhash:=hmac.New(sha256.New,key)
	// 给哈希对象添加数据
	myhash.Write(plainText)
	hashText:=myhash.Sum(nil)
	fmt.Println(hashText)
	return hashText
} 

// 验证消息认证码
func VeritfyHmac(plainText,key,hashText []byte) bool{
	// 创建哈希接口，需要指定哈希算法，和秘钥
	myhash:=hmac.New(sha256.New,key)
	// 给哈希对象添加数据
	myhash.Write(plainText)
	hmac1:=myhash.Sum(nil)
	// 比较两个散列值
	return hmac.Equal(hashText,hmac1)
}
func main() {
	src:=[]byte("这是一段消息体：你好，我是黄晓明")
	key:=[]byte("123456")
	hmacByte:=GenerateHamc(src,key)
	
	bl:=VeritfyHmac(src,key,hmacByte)
	fmt.Println(bl)
}