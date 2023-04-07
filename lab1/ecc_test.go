package main

import (
	"bytes"
	"math/big"
	"math/rand"
	"strconv"
	"testing"
)

func TestVerify(t *testing.T) {
	seckey, _ := new(big.Int).SetString("98710109312903829463955452893860670313169635532307033914348312391395047123341", 10)
	pubkey := GeneratePublicKey(seckey)
	ecc := MyECC{}
	s, _ := new(big.Int).SetString("28393339976117306978678592745526426404546994362727178347948902743375118960468", 10)
	r, _ := new(big.Int).SetString("33356462068120276508177129168229874346128066670085114831908640673373944866419", 10)
	sign := Signature{
		s: s,
		r: r,
	}
	msg := []byte("Dear Ms. Smith,I'm Li Hua, chairman of the Students' Union of Yucai Middle School, which is close to your university. I'm writing to invite you to be a judge at our English speech contest to be held in our school on June 15. It will start at 2:00 pm and last for about three hours. Ten students will deliver their speeches on the given topic \"Man and Nature\". We hope that you will accept our invitation if it is convenient for you. Please call me at 44876655 if you have any questions.I am looking forward to your reply.With best wishes,Li Hua")
	if !ecc.VerifySignature(msg, &sign, pubkey) {
		t.Error("Verify fail!")
	}
}

func TestVerifyError(t *testing.T) {
	seckey, _ := new(big.Int).SetString("98710109312903829463955452893860670313169635532307033914348312391395047123341", 10)
	pubkey := GeneratePublicKey(seckey)
	seckey1, _ := new(big.Int).SetString("98710109312903829463955452893860670313169635532307033914348312391395047123340", 10)
	pubkey1 := GeneratePublicKey(seckey1)
	ecc := MyECC{}
	s, _ := new(big.Int).SetString("28393339976117306978678592745526426404546994362727178347948902743375118960468", 10)
	r, _ := new(big.Int).SetString("33356462068120276508177129168229874346128066670085114831908640673373944866419", 10)
	s1, _ := new(big.Int).SetString("28393339976117306978678592745526426404546994362727178347948902743375118960469", 10)
	r1, _ := new(big.Int).SetString("33356462068120276508177129168229874346128066670085114831908640673373944866418", 10)
	sign := Signature{
		s: s,
		r: r,
	}
	sign1 := Signature{
		s: s1,
		r: r,
	}
	sign2 := Signature{
		s: s,
		r: r1,
	}
	msg := []byte("ear Ms. Smith,I'm Li Hua, chairman of the Students' Union of Yucai Middle School, which is close to your university. I'm writing to invite you to be a judge at our English speech contest to be held in our school on June 15. It will start at 2:00 pm and last for about three hours. Ten students will deliver their speeches on the given topic \"Man and Nature\". We hope that you will accept our invitation if it is convenient for you. Please call me at 44876655 if you have any questions.I am looking forward to your reply.With best wishes,Li Hua")
	if ecc.VerifySignature(msg, &sign, pubkey) {
		t.Error("Verify fail!")
	}
	msg2 := []byte("ustc")
	if ecc.VerifySignature(msg2, &sign, pubkey) {
		t.Error("Verify fail!")
	}
	if ecc.VerifySignature(msg, &sign1, pubkey) {
		t.Error("Verify fail!")
	}
	if ecc.VerifySignature(msg, &sign2, pubkey) {
		t.Error("Verify fail!")
	}
	if ecc.VerifySignature(msg, &sign, pubkey1) {
		t.Error("Verify fail!")
	}
}

func TestSign(t *testing.T) {
	seckey, _ := NewPrivateKey()
	pubkey := GeneratePublicKey(seckey)
	//seckey1, _ := NewPrivateKey()
	//pubkey1 := GeneratePublicKey(seckey1)
	ecc := MyECC{}

	msg := []byte("University of Science and Technology of China is one of the best universities in China")
	//msg1 := []byte("The future of blockchain technology is bright.")
	sign, err := ecc.Sign(msg, seckey)
	if err != nil {
		t.Error("Sign fail!")
	}
	if !ecc.VerifySignature(msg, sign, pubkey) {
		t.Error("Sign fail!")
	}
}

func TestSignRandom(t *testing.T) {
	seckey, _ := NewPrivateKey()
	pubkey := GeneratePublicKey(seckey)
	ecc := MyECC{}

	msg1 := new(bytes.Buffer)
	for i := 0; i < 10; i++ {
		msg1.WriteString(strconv.Itoa(rand.Intn(100)))
	}
	msg1.WriteString("hello!")
	sign1, err1 := ecc.Sign(msg1.Bytes(), seckey)
	if err1 != nil {
		t.Error("Sign fail!")
	}
	if !ecc.VerifySignature(msg1.Bytes(), sign1, pubkey) {
		t.Error("Sign fail!")
	}
}
