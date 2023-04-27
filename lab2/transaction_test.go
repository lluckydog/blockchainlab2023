package main

import (
	"testing"
	//"encoding/hex"
	//"github.com/stretchr/testify/assert"
)

func TestIsCoinbase(t *testing.T) {
	wallet, err := NewWallet()
	if err != nil {
		t.Errorf("new wallet error!")
	}
	address := wallet.GetAddress()
	coinbase := NewCoinbaseTx(address, address)
	if !coinbase.IsCoinBase() {
		t.Errorf("IsCoinbase is incorrect!")
	}
	txin := TXInput{[]byte{}, -1, nil, []byte(address)}
	oldcoinbasevin := append([]TXInput{}, coinbase.Vin...)
	newcoinbasevin := append(coinbase.Vin, txin)
	coinbase.Vin = newcoinbasevin
	if coinbase.IsCoinBase() {
		t.Errorf("IsCoinbase is incorrect!")
	}
	coinbase.Vin = oldcoinbasevin
	oldTxid := append(coinbase.Vin[0].Txid, []byte{}...)
	newTxid := append(coinbase.Vin[0].Txid, []byte("ok")...)
	coinbase.Vin[0].Txid = newTxid
	if coinbase.IsCoinBase() {
		t.Errorf("IsCoinbase is incorrect!")
	}
	coinbase.Vin[0].Txid = oldTxid
	coinbase.Vin[0].Vout = 0
	if coinbase.IsCoinBase() {
		t.Errorf("IsCoinbase is incorrect!")
	}
	coinbase.Vin[0].Vout = -1
	if !coinbase.IsCoinBase() {
		t.Errorf("IsCoinbase is incorrect!")
	}
}
