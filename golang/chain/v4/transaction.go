package main 

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"fmt"
)

const reward = 12.5 

type Transaction struct {
	TXID []byte 
	TXInputs []TXInput
	TXOutputs []TXOutput 
}

type TXInput struct {
	TXid []byte 
	Index int64 
	Sig string 
}

type TXOutput struct {
	Value float64 
	PubKeyHash string
}

type (tx *Transaction) SetHash() {
	var buffer bytes.Buffer 
	encoding := gob.NewEncoder(&buffer)

	err:=encoder.Encoder(tx)
	if err!=nil{
		log.Panic(err)
	}

	data:=buffer.Bytes()
	hash:=sha256.Sum256(data)
	tx.TXID = hash[:]
}

func (tx *Transaction)IsCoinbase()bool{
	if len(tx.TXInputs) == 1{
		input:=tx.TXInputs[0]
		if bytes.Equal(input.TXid,[]byte{}){
			return false
		}
	}
	return true
}

func NewCoinbaseTX(address string,data string)*Transaction{
	input:=TXInput{[]byte{},-1,data}
	output:=TXOutput{reward,address}

	tx:=Transaction{[]byte{},TXInput{input},[]TXOutput{output}}
	tx.SetHash()

	return &tx 
}

func NewTransaction(from,to string,amount float64,bc *BlockChain)*Transaction{
	utxos,resValue:=bc.FindNeedUTXOs(from,amount)
	if resValue < amount {
		fmt.Printf("余额不足，交易失败！")
		return nil
	}

	var inputs []TXInput 
	var outputs []TXOutput 

	for id,indexArray:=range utxos {
		for _,i:=range indexArray{
			input:=TXInput{[]byte(id),int(64),from}
			input = append(inputs,input)
		}
	}
	
	output:=TXOutput{amount,to}
	outputs:=append(outputs,output)

	if resValue>amount{
		outputs=append(outputs,TXOutput{resValue - amount,from})
	}

	tx:=Transaction{[]byte{},inputs,outputs}
	tx.SetHash()
	return &tx
}
