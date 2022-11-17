package main

import (
	"encoding/json"
	"fmt"

	"github.com/duck-dynasty/duck-blocks/accounts"
	"github.com/duck-dynasty/duck-blocks/blockchain"
)

func main() {
	fmt.Println("--------------- USER ACCOUNT TESTS ---------------")
	u1 := accounts.NewUserAccount()
	fmt.Println(u1.JSON())

	u2 := accounts.NewUserAccount()
	fmt.Println(u2.JSON())

	u3 := accounts.NewUserAccount()
	fmt.Println(u3.JSON())

	u4 := accounts.NewUserAccount()
	fmt.Println(u4.JSON())
	fmt.Println("--------------- USER ACCOUNT TESTS ---------------")
	fmt.Println("")

	fmt.Println("------------- USER TRANSACTION TESTS -------------")
	t1 := u1.CreateTransaction(u2.Address, 50)
	fmt.Println(t1.GetJSON())
	fmt.Println(t1.VerifyTransaction())

	t2 := u2.CreateTransaction(u4.Address, 20)
	fmt.Println(t2.GetJSON())
	fmt.Println(t2.VerifyTransaction())

	t3 := u3.CreateTransaction(u1.Address, 10)
	fmt.Println(t3.GetJSON())
	fmt.Println(t3.VerifyTransaction())

	t4 := u3.CreateTransaction(u4.Address, 15)
	fmt.Println(t4.GetJSON())
	fmt.Println(t4.VerifyTransaction())
	fmt.Println("------------- USER TRANSACTION TESTS -------------")
	fmt.Println("")

	s1 := accounts.AccountState{Address: u1.Address, Value: 100, Code: nil, Data: nil, Nonce: 0}
	s2 := accounts.AccountState{Address: u2.Address, Value: 100, Code: nil, Data: nil, Nonce: 0}
	s3 := accounts.AccountState{Address: u3.Address, Value: 100, Code: nil, Data: nil, Nonce: 0}

	bc := blockchain.NewBlockchain()
	bc.Chain[0].AccountStates[u1.Address] = s1
	bc.Chain[0].AccountStates[u2.Address] = s2
	bc.Chain[0].AccountStates[u3.Address] = s3
	bc.CurrentStates[u1.Address] = s1
	bc.CurrentStates[u2.Address] = s2
	bc.CurrentStates[u3.Address] = s3
	bc.NewTransaction(t1)
	bc.NewTransaction(t2)
	bc.NewTransaction(t3)
	bc.NewTransaction(t4)

	value, _ := json.MarshalIndent(bc, "", "    ")
	fmt.Println(string(value))
	/*
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

		fmt.Println("-------------------------------------")
		fmt.Println("Private Key: ", privateKey)
		fmt.Println("Public key: ", publicKey)
		fmt.Println("-------------------------------------")

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
	*/
}
