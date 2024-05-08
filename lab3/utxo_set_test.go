package main

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	//"encoding/hex"
	//"fmt"
	"testing"
)

func TestFindUnspentOutputs(t *testing.T) {
	//address, _ := hex.DecodeString("65376a4267326f6765535244505370584b77664e343948574d615479734a6b7338")
	bc := NewBlockchain()
	defer bc.Close()
	var utxoset UTXOSet
	utxoset.Blockchain = bc
	pkh := []byte{
		0x65, 0xA9, 0x85, 0x35, 0xA5, 0xAD, 0xD0, 0x68,
		0x02, 0x0C, 0xE2, 0x52, 0x47, 0x45, 0x77, 0xEB,
		0x60, 0xFE, 0x96, 0x92,
	}
	i, j := utxoset.FindUnspentOutputs(pkh, 200000)
	fmt.Printf("%v", j)
	assert.Equal(t, 209999, i, "findunspentoutputs fails!")
	assert.Equal(t, j["5a277a0a503382ce49aee4aeb74ad61651911d2eb44dde01a18e5d78776954ed"], []int{1}, "findunspentoutputs fails!")
}
