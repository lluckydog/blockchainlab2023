package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	data := []byte("data")
	var lastHash [32]byte
	for i := 0; i < 32; i++ {
		lastHash[i] = byte(32 - i)
	}
	toaddr := []byte("toadress")
	tx := NewCoinbaseTx(toaddr, data)
	var txs Transactions
	txs = append(txs, tx)
	txs = append(txs, tx)
	blkheader := BlkHeader{
		Version:       int64(100),
		PrevBlockHash: lastHash,
		MerkleRoot:    txs.CalculateHash(),
		Bits:          targetBits,
		Timestamp:     int64(20240508),
	}
	blkbody := BlkBody{txs}
	block := &Block{&blkheader, &blkbody}
	pow := NewProofOfWork(block)
	target := big.NewInt(1)
	target.Lsh(target, uint(246))
	pow.target = target
	nonce, _ := pow.Run()
	block.Header.Nonce = nonce
	assert.Equal(t, nonce, int64(121), "pow run fail!")
}

func TestValidate(t *testing.T) {
	data := []byte("data")
	var lastHash [32]byte
	for i := 0; i < 32; i++ {
		lastHash[i] = byte(i)
	}
	toaddr := []byte("toadress")
	tx := NewCoinbaseTx(toaddr, data)
	var txs Transactions
	txs = append(txs, tx)
	txs = append(txs, tx)
	blkheader := BlkHeader{
		Version:       int64(1),
		PrevBlockHash: lastHash,
		MerkleRoot:    txs.CalculateHash(),
		Bits:          targetBits,
		Timestamp:     int64(20240508),
	}
	blkbody := BlkBody{txs}
	block := &Block{&blkheader, &blkbody}
	pow := NewProofOfWork(block)
	target := big.NewInt(1)
	target.Lsh(target, uint(246))
	pow.target = target

	nonce, _ := pow.Run()
	pow.block.Header.Nonce = nonce
	block.Header.Nonce = int64(326)
	if pow.Validate() {
		t.Error("pow validate fail!")
	}
	block.Header.Nonce = int64(336)
	if !pow.Validate() {
		t.Error("pow validate fail!")
	}

}
