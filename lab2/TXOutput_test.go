package main

import (
	"testing"
	//"encoding/hex"
	"github.com/stretchr/testify/assert"
)
func TestLock(t *testing.T) {
	address := []byte("Think of this string as an address")
	var txoutput TXOutput
	txoutput.Value = 1
	txoutput.PubKeyHash = nil
	txoutput.Lock(address)
	realph := []byte("ug3GZyB8HTJawvqRYhhERD9k5Apdm6DhckYjTyzcJK")
	assert.Equal(t, realph, txoutput.PubKeyHash, "Lock is incorrect!")
}