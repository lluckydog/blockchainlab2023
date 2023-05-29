package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMineBlock(t *testing.T) {
	//address, _ := hex.DecodeString("65376a4267326f6765535244505370584b77664e343948574d615479734a6b7338")
	toaddr, _ := hex.DecodeString("20277e9c9fa0838769ec353bd8c390c6aec277d6d5e46580b7310ac13b13fcba57")
	bc := NewBlockchain()
	defer bc.Close()
	data := []byte("")
	tx := NewCoinbaseTx(toaddr, data)
	txs := make([]*Transaction, 0)
	txs = append(txs, tx)
	toaddr, _ = hex.DecodeString("575358554b387950714e55696159566f6f6371767163766d533464457539574836")
	tx = NewCoinbaseTx(toaddr, data)
	txs = append(txs, tx)
	block := bc.MineBlock(txs)
	assert.Equal(t, int64(1), block.Header.Version, "mine block fail!")
	mk := [32]uint8([32]uint8{0x93, 0xc5, 0x36, 0x68, 0x64, 0x20, 0x52, 0xda, 0x61, 0x9, 0x5d, 0x16, 0xd4, 0xdb, 0x92, 0x9d, 0x1e, 0x15, 0x3d, 0x56, 0x62, 0x2a, 0x9c, 0xa4, 0x39, 0x38, 0x49, 0x14, 0x36, 0x5f, 0x12, 0xc9})
	assert.Equal(t, mk, block.Header.MerkleRoot, "mine block fail!")
	assert.Equal(t, bc.tip, block.CalCulHash(), "mine block fail!")
	pow := NewProofOfWork(block)
	if !pow.Validate() {
		t.Error("pow validate fail!")
	}
}

func TestFindUTXO(t *testing.T) {
	//address, _ := hex.DecodeString("65376a4267326f6765535244505370584b77664e343948574d615479734a6b7338")
	bc := NewBlockchain()
	defer bc.Close()
	r := bc.FindUTXO()
	for k, v := range r {
		fmt.Printf("%s %d\n", k, v)
	}
	k := "0fe552774978000c3cc084b832d44172f31371c9a27f07a6ce70044b04f055f3"
	v, ok := r[k]
	if !ok || v.Outputs[0].Value != 210000 {
		t.Error("find utxo fail!")
	}
	pk := []byte{0x51, 0x54, 0x67, 0x79, 0x41, 0x54, 0x33, 0x48, 0x54, 0x47, 0x52, 0x66, 0x46, 0x33, 0x67, 0x4a, 0x48, 0x51, 0x43, 0x59, 0x37, 0x6d, 0x6d, 0x56, 0x78, 0x6d, 0x36, 0x4e, 0x75, 0x36, 0x55, 0x65, 0x4d}
	assert.Equal(t, pk, v.Outputs[0].PubKeyHash, "find utxo fail!")
	k2 := "b294a1dec98fbf06ace81ae2d2f0e18a562578573fb6047ce616548d294ffebd"
	v2, ok2 := r[k2]
	pkh1 := []byte{0x73, 0x5a, 0x91, 0x6c, 0x67, 0xb8, 0x26, 0xcc, 0x9c, 0x92, 0x69, 0x71, 0x61, 0x51, 0xf4, 0x98, 0x51, 0x7d, 0x24, 0xe4}
	pkh2 := []byte{0x1, 0x57, 0x12, 0x68, 0xd0, 0x4e, 0x43, 0x3c, 0x62, 0xac, 0xf1, 0xf8, 0x38, 0xf3, 0xc4, 0x44, 0xb9, 0xb6, 0xda, 0xbc}
	if !ok2 {
		t.Error("find utxo fail!")
	}
	assert.Equal(t, pkh1, v2.Outputs[0].PubKeyHash, "find utxo fail!")
	assert.Equal(t, pkh2, v2.Outputs[1].PubKeyHash, "find utxo fail!")
	assert.Equal(t, 1, v2.Outputs[0].Value, "find utxo fail!")
	assert.Equal(t, 209999, v2.Outputs[1].Value, "find utxo fail!")
}
