package main 

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"fmt"
)

const reward = 50 

type Transaction struct {
	TXID []byte 
	TXInput []TXInput 
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