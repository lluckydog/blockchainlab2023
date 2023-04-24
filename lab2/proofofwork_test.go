package main

import (
	"strconv"
	"fmt"
	"testing"
	"math/big"
	"time"
	"crypto/sha256"
)

func TestRun(t *testing.T)  {
	var datas [][]byte
	for i:=0;i<5;i++ {
		datas = append(datas, []byte("ok"+strconv.Itoa(i)))
	}
	lastHash := []byte("just test ww")
	block := &Block{time.Now().Unix(), datas, lastHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	target := big.NewInt(1)
	target.Lsh(target, uint(246))
	pow.target = target
	nonce, _ := pow.Run()
	fmt.Printf("Nonce : %d\n", nonce)
	block.Nonce = nonce
	thash := sha256.Sum256(pow.block.Serialize())
	block.Hash = thash[:]
	
	if new(big.Int).SetBytes(thash[:]).Cmp(pow.target) >= 0 {
		t.Error("pow run fail!")
	}
}

func TestValidate(t *testing.T) {
	var datas [][]byte
	for i:=0;i<5;i++ {
		datas = append(datas, []byte("ok"+strconv.Itoa(i)))
	}
	lastHash := []byte("just test")
	block := &Block{1682245264, datas, lastHash, []byte{}, 1984}
	pow := NewProofOfWork(block)
	target := big.NewInt(1)
	target.Lsh(target, uint(246))
	pow.target = target
	thash := sha256.Sum256(pow.block.Serialize())
	block.Hash = thash[:]
	if !pow.Validate() {
		t.Error("pow validate fail!")
	}

	fblock := &Block{1682245264, datas, lastHash, []byte{}, 1980}
	fpow := NewProofOfWork(fblock)
	fpow.target = target
	fhash := sha256.Sum256(fpow.block.Serialize())
	fblock.Hash = fhash[:]
	if fpow.Validate() {
		t.Error("pow validate fail!")
	}
}