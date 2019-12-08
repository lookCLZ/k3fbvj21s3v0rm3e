package main

import (
	// Elliptic Curve Digital Signature Algorithm
	// 椭圆曲线数字签名算法

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"crypto/rand"
	"crypto/sha1"

	"encoding/pem"
	"math/big"

	"os"
	"fmt"
)

// 生成密钥对 使用ecdsa生成密钥对
func GenerateEccKey() {
	// 生成私钥
	privateKey,_:=ecdsa.GenerateKey(elliptic.P521(),rand.Reader)
	// 使用x509标准序列化
	derText,_:=x509.MarshalECPrivateKey(privateKey)
	block:=pem.Block {
		Type:"椭圆曲线加密算法私钥",
		Bytes:derText,
	}
	file,_:=os.Create("eccPrivate.pem")
	// 将block写入文件
	pem.Encode(file,&block)
	file.Close()

	// 获取公钥
	publicKey:=privateKey.PublicKey
	derText,_=x509.MarshalPKIXPublicKey(&publicKey)

	// 将得到的切片放入block里面
	block=pem.Block{
		Type:"ecdsa public key",
		Bytes:derText,
	}
	// 使用pem编码
	file,_=os.Create("eccPublic.pem")
	pem.Encode(file,&block)
	file.Close()
}

// 签名
func EccSignature(plainText []byte,privName string) (rText,sText []byte) {
	// 获取私钥
	file,_:=os.Open(privName)
	fileInfo,_:=file.Stat()
	buf:=make([]byte,fileInfo.Size())
	file.Read(buf)
	file.Close()
	// 解码成结构体
	block,_:=pem.Decode(buf)
	// 使用x509,反序列化
	privateKey,_:=x509.ParseECPrivateKey(block.Bytes)
	// 算出hash值
	hashText:=sha1.Sum(plainText)
	// 进行数字签名
	r,s,_:=ecdsa.Sign(rand.Reader,privateKey,hashText[:])
	// 对r,s内存中的数据进行格式化 -> []byte
	rText,_=r.MarshalText()
	sText,_=s.MarshalText()

	return
}

// ecc签名认证
func EccVerify(plainText,rText,sText []byte,pubFile string) bool {
	// 获取公钥
	file,_:=os.Open(pubFile)
	fileInfo,_:=file.Stat()
	buf:=make([]byte,fileInfo.Size())
	file.Read(buf)
	file.Close()
	block,_:=pem.Decode(buf)
	pubInterface,_:=x509.ParsePKIXPublicKey(block.Bytes)
	pubKey:=pubInterface.(*ecdsa.PublicKey)
	// 获取hash值
	hashText:=sha1.Sum(plainText)
	// 将rText,sText -> int
	var r,s big.Int
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)
	// 认证
	bl:=ecdsa.Verify(pubKey,hashText[:],&r,&s)
	return bl
}
func main() {
	GenerateEccKey()
	src:=[]byte("大家好，我是张三丰")
	rText,sText:=EccSignature(src,"eccPrivate.pem")
	bl:=EccVerify(src,rText,sText,"eccPublic.pem")
	fmt.Println(bl)
}

