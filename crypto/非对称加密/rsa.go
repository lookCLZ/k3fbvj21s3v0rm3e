package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/sha256"

	"encoding/pem"
	"encoding/hex"

	"os"
	"fmt"
)

func GenerateRsaKey(keySize int) {
	// 使用RSA中的GenerateKey生成私钥
	privateKey,err:=rsa.GenerateKey(rand.Reader,keySize)
	if err!=nil{
		panic(err)
	}
	// 通过x509标准 序列化 编码字符串
	derText:=x509.MarshalPKCS1PrivateKey(privateKey)
	// 要组织一个pem.Block(base64编码)
	block:=pem.Block{
		Type:"这是一个私钥",
		Bytes:derText,
	}
	// pem编码
	file,err:=os.Create("private.pem")
	if err!=nil{
		panic(err)
	}
	// 写入文件
	pem.Encode(file,&block)
	file.Close()

	// 创建公钥
	// 从私钥中提取公钥
	publicKey:=privateKey.PublicKey
	// 使用x509标准序列化
	derstream,err:=x509.MarshalPKIXPublicKey(&publicKey)
	if err!=nil{
		panic(err)
	}
	// 将数据放到pem.Block中
	block=pem.Block{
		Type:"rsa public key",
		Bytes:derstream,
	}
	// pem编码
	file,err=os.Create("public.pem")
	if err!=nil{
		panic(err)
	}
	pem.Encode(file,&block)
	file.Close()
}

// RSA 公钥加密
func RSAEncrypt(plainText []byte,fileName string) []byte{
	// 打开文件，读取内容
	file,err:=os.Open(fileName)
	if err!=nil{
		panic(err)
	}
	// 获取文件状态信息
	fileInfo,err:=file.Stat()
	if err!=nil{
		panic(err)
	}
	buf:=make([]byte,fileInfo.Size())
	file.Read(buf)
	file.Close()
	// pem解码到 block 结构体
	block,_:=pem.Decode(buf)
	// 获取公钥
	pubInterface,err:=x509.ParsePKIXPublicKey(block.Bytes)
	// 断言类型转换
	pubKey:=pubInterface.(*rsa.PublicKey)
	// 使用公钥进行加密
	cipherText,err:=rsa.EncryptPKCS1v15(rand.Reader,pubKey,plainText)
	if err!=nil{
		panic(err)
	}
	return cipherText
}

// RSA 使用私钥解密
func RSADecrypt(cipherText []byte,fileName string)[]byte{
	// 打开文件，读取内容
	file,err:=os.Open(fileName)
	if err!=nil{
		panic(err)
	}
	fileInfo,_:=file.Stat()
	buf:=make([]byte,fileInfo.Size())
	file.Read(buf)
	file.Close()

	// pem解码到结构体 block
	block,_:=pem.Decode(buf)
	// 获取私钥
	privKey,err:=x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	// 使用私钥解密
	plainText,err:=rsa.DecryptPKCS1v15(rand.Reader,privKey,cipherText)
	if err!=nil{
		panic(err)
	}
	return plainText
}

func Hash() {
	// 建立hash接口对象
	hash:=sha256.New()
	// 添加数据
	src:=[]byte("天龙八部")
	// hash.Write(src)
	// 
	res:=hash.Sum(src)
	fmt.Println(res)
	str:=hex.EncodeToString(res)
	fmt.Println(str)

}

func main() {
	GenerateRsaKey(4096)
	src:=[]byte("我是好迪，大家好，才是真的好！")
	cipherText:=RSAEncrypt(src,"public.pem")
	plainText:=RSADecrypt(cipherText,"private.pem")
	fmt.Println(string(plainText))

	Hash()
	
}