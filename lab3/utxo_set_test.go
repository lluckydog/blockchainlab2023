package main

import (
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
	pkh := []byte{0x1, 0x57, 0x12, 0x68, 0xd0, 0x4e, 0x43, 0x3c, 0x62, 0xac, 0xf1, 0xf8, 0x38, 0xf3, 0xc4, 0x44, 0xb9, 0xb6, 0xda, 0xbc}
	i, j := utxoset.FindUnspentOutputs(pkh, 200000)
	//fmt.Printf("%d\n", i)
	assert.Equal(t, 209999, i, "findunspentoutputs fails!")
	assert.Equal(t, j["b294a1dec98fbf06ace81ae2d2f0e18a562578573fb6047ce616548d294ffebd"], []int{1}, "findunspentoutputs fails!")
}
