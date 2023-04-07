package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

var (
	s256    *secp256k1.BitCurve = secp256k1.S256()
	P, _                        = new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 0)
	N, _                        = new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 0)
	B, _                        = new(big.Int).SetString("0x0000000000000000000000000000000000000000000000000000000000000007", 0)
	Gx, _                       = new(big.Int).SetString("0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 0)
	Gy, _                       = new(big.Int).SetString("0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 0)
	BitSize                     = 256
	G                           = &Point{Gx, Gy}
)

type Point struct {
	X *big.Int
	Y *big.Int
}

type Signature struct {
	s *big.Int
	r *big.Int
}

type ECC interface {
	Sign(msg []byte, secKey *big.Int) (*Signature, error)
	VerifySignature(msg []byte, signature *Signature, pubkey *Point) bool
}

type MyECC struct {
}

func NewPrivateKey() (*big.Int, error) {
	k, err := newRand()
	if err != nil {
		return nil, err
	}
	if err := checkBigIntSize(k); err != nil {
		return nil, fmt.Errorf("k error: %s", err)
	}

	return k, nil
}

func GeneratePublicKey(secKey *big.Int) *Point {
	return Multi(G, secKey)
}

func (ecc *MyECC) Sign(msg []byte, secKey *big.Int) (*Signature, error) {
	return nil, nil
}

// >>> point = S256Point(px, py)
// >>> s_inv = pow(s, N-2, N)  ❶
// >>> u = z * s_inv % N  ❷
// >>> v = r * s_inv % N  ❸
// >>> print((u*G + v*point).x.num == r)
func (ecc *MyECC) VerifySignature(msg []byte, signature *Signature, pubkey *Point) bool {
	return true
}

func main() {
	seckey, err := NewPrivateKey()
	if err != nil {
		fmt.Println("error!")
	}
	pubkey := GeneratePublicKey(seckey)

	ecc := MyECC{}
	msg := []byte("test1")
	msg2 := []byte("test2")

	sign, err := ecc.Sign(msg, seckey)
	if err != nil {
		fmt.Printf("err %v\n", err)
		return
	}

	fmt.Printf("verify %v\n", ecc.VerifySignature(msg, sign, pubkey))
	fmt.Printf("verify %v\n", ecc.VerifySignature(msg2, sign, pubkey))

}
