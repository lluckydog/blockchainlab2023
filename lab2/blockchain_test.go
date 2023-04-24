package main

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
)

func TestAddBlock(t *testing.T) {
	blockchain := NewBlockchain()
	var stack []*Block
	var datas [][]byte
	for i := 0; i < 5; i++ {
		datas = append(datas, []byte("ok"+strconv.Itoa(5-i)))
	}
	err := blockchain.AddBlock(datas)
	if err != nil {
		t.Error("add block fail!")
	}
	for i := 0; i < 5; i++ {
		datas = append(datas, []byte("ok"+strconv.Itoa(i)))
	}
	err = blockchain.AddBlock(datas)
	if err != nil {
		t.Error("add block fail!")
	}
	bcin := blockchain.Iterator()

	nowblock := bcin.Next()
	stack = append(stack, nowblock)
	assert.Equal(t, datas[9], nowblock.Data[9], "add block fail!(data)")
	pow := NewProofOfWork(nowblock)
	if new(big.Int).SetBytes(pow.block.Hash).Cmp(pow.target) >= 0 {
		t.Error("add block fail!(nonce)")
	}
	nowblock = bcin.Next()
	stack = append(stack, nowblock)
	genehash := nowblock.Hash
	assert.Equal(t, datas[4], nowblock.Data[4], "add block fail!(data)")
	var block *Block
	i := blockchain.Iterator()
	_ = i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	assert.Equal(t, block.PrevBlockHash, genehash, "add block fail!(preblockhash)")
	nowblock = bcin.Next()
	stack = append(stack, nowblock)
	assert.Equal(t, stack[2].Hash, stack[1].PrevBlockHash, "add block fail!(preblockhash)")
	assert.Equal(t, stack[1].Hash, stack[0].PrevBlockHash, "add block fail!(preblockhash)")
}
