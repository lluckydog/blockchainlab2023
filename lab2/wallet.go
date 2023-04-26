package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x01)
const checkSumlen = 4
const walletFile = "wallet.dat"

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallet() (*Wallet, error) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	public := private.PublicKey

	return &Wallet{
		PrivateKey: *private,
		PublicKey:  append(public.X.Bytes(), public.Y.Bytes()...),
	}, nil
}

func HashPublicKey(pubKey []byte) []byte {
	publicSha256 := sha256.Sum256(pubKey)

	Hasher := ripemd160.New()
	_, err := Hasher.Write(publicSha256[:])
	if err != nil {
		log.Panic("write hash error")
	}

	return Hasher.Sum(nil)
}

func (w *Wallet) GetAddress() []byte {
	return nil
}

// NewWallets creates Wallets and fills it from a file if it exists
func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile()

	return &wallets, err
}

// CreateWallet adds a Wallet to Wallets
func (ws *Wallets) CreateWallet() string {
	wallet, err := NewWallet()
	if err != nil {
		log.Panic("create wallet fail")
	}
	address := wallet.GetAddress()

	ws.Wallets[hex.EncodeToString(address)] = wallet

	return hex.EncodeToString(address)
}

// GetAddresses returns an array of addresses stored in the wallet file
func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

// GetWallet returns a Wallet by its address
func (ws Wallets) GetWallet(address []byte) Wallet {
	return *ws.Wallets[hex.EncodeToString(address)]
}

// LoadFromFile loads wallets from the file
func (ws *Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	ws.Wallets = wallets.Wallets

	return nil
}

// SaveToFile saves wallets to a file
func (ws Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
