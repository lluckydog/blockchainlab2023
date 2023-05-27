package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func DeSerializeRS(signature []byte) (*big.Int, *big.Int) {
	sigLen := len(signature)
	r := new(big.Int).SetBytes(signature[:sigLen/2])
	s := new(big.Int).SetBytes(signature[sigLen/2:])

	return r, s
}
