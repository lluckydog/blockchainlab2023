package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/btcsuite/btcutil/base58"
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

func CheckSum(payload []byte) []byte {
	inter := sha256.Sum256(payload)
	res := sha256.Sum256(inter[:])
	return res[:checkSumlen]
}

func (w *Wallet) GetAddress() []byte {
	pubKeyHash := HashPublicKey(w.PublicKey)

	payload := append([]byte{version}, pubKeyHash...)
	checksum := CheckSum(payload)
	payload = append(payload, checksum...)
	address := []byte(base58.Encode(payload))
	return address
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

	err = json.Unmarshal(fileContent, ws)
	if err != nil {
		log.Panic(err)
	}

	return nil
}

// SaveToFile saves wallets to a file
func (ws *Wallets) SaveToFile() {
	jsonData, err := json.Marshal(ws)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, jsonData, 0644)
	if err != nil {
		log.Panic(err)
	}
}

func (w *Wallet) MarshalJSON() ([]byte, error) {
	// 私钥的D值和公钥都以Base64编码存储
	keyData := map[string]string{
		"D":         base64.StdEncoding.EncodeToString(w.PrivateKey.D.Bytes()),
		"PublicKey": base64.StdEncoding.EncodeToString(w.PublicKey),
	}
	return json.Marshal(keyData)
}

func (w *Wallet) UnmarshalJSON(data []byte) error {
	keyData := make(map[string]string)
	if err := json.Unmarshal(data, &keyData); err != nil {
		return err
	}

	dBytes, err := base64.StdEncoding.DecodeString(keyData["D"])
	if err != nil {
		return err
	}
	publicKeyBytes, err := base64.StdEncoding.DecodeString(keyData["PublicKey"])
	if err != nil {
		return err
	}

	// 重构PrivateKey和PublicKey
	privateKey := new(ecdsa.PrivateKey)
	privateKey.D = new(big.Int).SetBytes(dBytes)
	privateKey.PublicKey.Curve = elliptic.P256()
	privateKey.PublicKey.X, privateKey.PublicKey.Y = elliptic.Unmarshal(elliptic.P256(), publicKeyBytes)

	w.PrivateKey = *privateKey
	w.PublicKey = publicKeyBytes

	return nil
}

// ValidateAddress check if address if valid
func ValidateAddress(address string) bool {
	data, _ := hex.DecodeString(address)

	pubKeyHash := base58.Decode(string(data))
	actualChecksum := pubKeyHash[len(pubKeyHash)-checkSumlen:]
	// version := pubKeyHash[0]
	// pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-checkSumlen]
	targetChecksum := CheckSum(pubKeyHash[:len(pubKeyHash)-checkSumlen])
	return bytes.Equal(actualChecksum, targetChecksum)
}
