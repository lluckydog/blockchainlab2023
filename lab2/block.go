package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

const curVersion = 1

// Block keeps block headers
type Block struct {
	Header *BlkHeader
	Body   *BlkBody
}

type BlkHeader struct {
	Version       int64
	PrevBlockHash []byte
	MerkleRoot    []byte
	Timestamp     int64
	Bits          int64
	Nonce         int64
}

type BlkBody struct {
	Transactions Transactions
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func (b *Block) SetNonce(nonce int64) {
	b.Header.Nonce = nonce
}

func (b *Block) GetTransactions() Transactions {
	return b.Body.Transactions
}

func (b *Block) GetPrevhash() []byte {
	return b.Header.PrevBlockHash
}

func (b *Block) CalCulHash() []byte {
	res := sha256.Sum256(b.Serialize())
	return res[:]
}

func NewBlkHeader(transactions Transactions, prevBlockHash []byte) *BlkHeader {
	return &BlkHeader{
		Version:       curVersion,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    transactions.CalculateHash(),
		Timestamp:     time.Now().Unix(),
		Bits:          targetBits,
		Nonce:         0,
	}
}

func NewBlkBody(transactions Transactions) *BlkBody {
	return &BlkBody{transactions}
}

// NewBlock creates and returns Block
func NewBlock(transactions Transactions, prevBlockHash []byte) *Block {
	head := NewBlkHeader(transactions, prevBlockHash)
	body := NewBlkBody(transactions)
	block := &Block{
		Header: head,
		Body:   body,
	}
	pow := NewProofOfWork(block)
	nonce, _ := pow.Run()

	block.SetNonce(nonce)

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coionbase *Transaction) *Block {
	return NewBlock([]*Transaction{coionbase}, []byte{})
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
