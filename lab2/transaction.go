package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
)

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type Transactions []*Transaction

const baseValue int = 210000

func NewCoinbaseTx(toAddr []byte, data []byte) *Transaction {
	txin := TXInput{[]byte{}, -1, nil, []byte(data)}
	txout := TXOutput{baseValue, toAddr}

	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}

// NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(from, to []byte, amount int, UTXOSet *UTXOSet) *Transaction {
	return nil
}

func (t *Transaction) IsCoinBase() bool {
	return true
}

func (t *Transaction) Serialize() []byte {
	var encode bytes.Buffer

	enc := gob.NewEncoder(&encode)
	err := enc.Encode(t)
	if err != nil {
		log.Panic(err)
	}

	return encode.Bytes()
}

func (t *Transaction) CalCulHash() []byte {
	txCopy := *t
	txCopy.ID = []byte{}
	val := sha256.Sum256(txCopy.Serialize())

	return val[:]
}

func (t *Transaction) SetID() {
	t.ID = t.CalCulHash()
}

func (t *Transaction) TrimmedCopy() *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	for _, vin := range t.Vin {
		inputs = append(inputs, TXInput{vin.Txid, vin.Vout, nil, nil})
	}

	for _, vout := range t.Vout {
		outputs = append(outputs, TXOutput{vout.Value, vout.PubKeyHash})
	}

	txCopy := Transaction{t.ID, inputs, outputs}

	return &txCopy
}

func (t *Transaction) Sign(privkey ecdsa.PrivateKey, prevTXs map[string]*Transaction) error {
	if t.IsCoinBase() {
		return nil
	}

	txCopy := t.TrimmedCopy()

	for i, vin := range txCopy.Vin {
		prevTX := prevTXs[hex.EncodeToString(vin.Txid)]
		txCopy.Vin[i].PubKey = prevTX.Vout[vin.Vout].PubKeyHash
		txCopy.SetID()

		r, s, err := ecdsa.Sign(rand.Reader, &privkey, txCopy.ID)
		if err != nil {
			return err
		}
		signature := append(r.Bytes(), s.Bytes()...)

		t.Vin[i].Signature = signature
	}
	return nil
}

func (t *Transaction) Verify(prevTXs map[string]*Transaction) bool {
	return true
}

func (txs Transactions) Serialize() [][]byte {
	var payload [][]byte
	for _, tx := range txs {
		payload = append(payload, tx.Serialize())
	}

	return payload
}

func (txs Transactions) CalculateHash() []byte {
	tree := NewMerkleTree(txs.Serialize())
	return tree.RootNode.Data
}
