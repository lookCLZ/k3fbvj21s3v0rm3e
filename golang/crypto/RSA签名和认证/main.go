package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// RSA签名
func SignatureRSA(plainText []byte,fileName string) []byte{
	// 打开私钥文件
	file,err:=os.Open(fileName)
	if err!=nil{
		panic(err)
	}
	// 读出私钥
	fileInfo,_:=file.Stat()
	buf:=make([]byte,fileInfo.Size())
	file.Read(buf)
	file.Close()

	// 解码
	block,_:=pem.Decode(buf)
	// 使用x509标准将私钥解析成结构体
	priKey,_:=x509.ParsePKCS1PrivateKey(block.Bytes)
	// 将消息生成散列值
	// hashText:=sha512.Sum512(plainText)

	myhash := sha512.New()
	//6. 给哈希对象添加数据
	myhash.Write(plainText)
	//7. 计算哈希值
	hashText := myhash.Sum(nil)

	// 使用rsa中的函数对散列值签名 (私钥，消息哈希)
	sigText,_:=rsa.SignPKCS1v15(rand.Reader,priKey,crypto.SHA512,hashText)

	return sigText
}

func VerifyRSA(plainText,sigText []byte,pubFileName string) bool{
	// 打开公钥文件
	file,_:=os.Open(pubFileName)
	info,_:=file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)
	file.Close()

	// 使用pem解码，得到block结构体变量
	block,_:=pem.Decode(buf)
	// 使用x509标准解码，将公钥解析出来
	pubInterface,_:=x509.ParsePKIXPublicKey(block.Bytes)
	pubKey:=pubInterface.(*rsa.PublicKey)
	// 计算消息的散列值
	hashText:=sha512.Sum512(plainText)

	// 认证签名
	err:=rsa.VerifyPKCS1v15(pubKey,crypto.SHA512,hashText[:],sigText)
	if err ==nil{
		return true
	}
	return false
}

func main() {
	src:=[]byte("哈哈，这是一段文本哦，😁")
	sigText:=SignatureRSA(src,"private.pem")
	bl:=VerifyRSA(src,sigText,"public.pem")

	fmt.Println(bl)
}