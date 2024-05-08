package main

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMineBlock(t *testing.T) {
	//address, _ := hex.DecodeString("65376a4267326f6765535244505370584b77664e343948574d615479734a6b7338")
	toaddr, _ := hex.DecodeString("5951674b7531543174736b3543636552614e4c6a756e77645741697555436f356f")
	bc := NewBlockchain()
	defer bc.Close()
	data := []byte("")
	tx := NewCoinbaseTx(toaddr, data)
	txs := make([]*Transaction, 0)
	txs = append(txs, tx)
	toaddr, _ = hex.DecodeString("6d51717833706a347a684a506a4267796d553471774c7433323274707868556a33")
	tx = NewCoinbaseTx(toaddr, data)
	txs = append(txs, tx)
	block := bc.MineBlock(txs)
	assert.Equal(t, int64(1), block.Header.Version, "mine block fail!")
	mk := [32]uint8([32]uint8{0x4d, 0x7d, 0x59, 0x57, 0x15, 0x43, 0x85, 0xe9, 0x4c, 0xe9, 0xb3, 0x3e, 0x11, 0x51, 0x2, 0x80, 0x51, 0xd0, 0xb4, 0xf0, 0xc4, 0xef, 0x3b, 0x8f, 0xc5, 0xb4, 0x73, 0xf, 0xae, 0x1c, 0x28, 0x14})
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
	k := "5a277a0a503382ce49aee4aeb74ad61651911d2eb44dde01a18e5d78776954ed"
	v, ok := r[k]
	if !ok {
		t.Error("find utxo fail!")
	}
	pk := []byte{
		0xe7, 0x27, 0xf6, 0x57, 0x61, 0xb9, 0x73, 0x61, 0x6f, 0xc7, 0x79, 0xac, 0x85, 0xfd, 0x2b, 0x40, 0x61, 0x59, 0x2, 0xfd,
	}

	pkh := []byte{
		0xe7, 0x27, 0xf6, 0x57, 0x61, 0xb9, 0x73, 0x61, 0x6f, 0xc7, 0x79, 0xac, 0x85, 0xfd, 0x2b, 0x40, 0x61, 0x59, 0x2, 0xfd,
	}
	assert.Equal(t, pk, v.Outputs[0].PubKeyHash, "find utxo fail!")
	assert.Equal(t, pkh, v.Outputs[0].PubKeyHash, "find utxo fail!")
	assert.Equal(t, 1, v.Outputs[0].Value, "find utxo fail!")
	assert.Equal(t, 209999, v.Outputs[1].Value, "find utxo fail!")
}
