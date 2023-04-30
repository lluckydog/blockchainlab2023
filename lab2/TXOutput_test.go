package main

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLock(t *testing.T) {
	address, _ := hex.DecodeString("65376a4267326f6765535244505370584b77664e343948574d615479734a6b7338")

	var txoutput TXOutput
	txoutput.Value = 1
	txoutput.PubKeyHash = nil
	txoutput.Lock(address)

	realph, _ := hex.DecodeString("97225cfb988e1a53474ef6ada09eea461b8047da")
	assert.Equal(t, realph, txoutput.PubKeyHash, "Lock is incorrect!")
}
