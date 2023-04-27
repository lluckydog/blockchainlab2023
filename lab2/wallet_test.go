package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddress(t *testing.T) {
	var wallet2 Wallet
	curve := elliptic.P256()
	x, _ := hex.DecodeString("ebd8e60ff9ed42ffe8ced4850d18e6c581b6f34f7440e095972dbd86d15e6872")
	y, _ := hex.DecodeString("57d892f6fa44c291d31ec82ce43fb60a0834081cf53aa54f7029bf3fb9cd6909")
	d, _ := hex.DecodeString("a1ad7be296f44c1d63158c8476bd180b050bebc0cdc002dbf13b57af7c607ef1")
	public := ecdsa.PublicKey{
		curve,
		new(big.Int).SetBytes(x),
		new(big.Int).SetBytes(y),
	}
	wallet2.PrivateKey = ecdsa.PrivateKey{
		public,
		new(big.Int).SetBytes(d),
	}
	wallet2.PublicKey = append(public.X.Bytes(), public.Y.Bytes()...)

	address2, _ := hex.DecodeString("637361657a6e54346b363944414e617642466f46664767477673574872746f5669")
	assert.Equal(t,
		hex.EncodeToString(wallet2.GetAddress()),
		hex.EncodeToString(address2),
		"Getting wallet address fails!")

}
