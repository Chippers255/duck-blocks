package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	transactions "github.com/duck-dynasty/duck-blocks/transactions"
)

func main() {
	tv := transactions.DuckartParams{Amount: 1000000000000000000, From: "Tom", MaxFee: 12345, Nonce: 1, To: "Tim", Type: 127}
	t1 := transactions.DuckatTransaction{Params: &tv}

	min_value := t1.GetJSON()
	fmt.Println(min_value)
	fmt.Println("")

	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	t1.SignTransaction(privateKey)
	min_value = t1.GetJSON()
	fmt.Println(min_value)
	fmt.Println("")

	msg := []byte(t1.GetJSON())

	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err = msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return
	}

	fmt.Println("signature verified")
	fmt.Println(publicKey)
}
