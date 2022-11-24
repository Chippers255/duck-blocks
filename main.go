package main

import (
	"encoding/json"
	"fmt"

	"github.com/duck-dynasty/duck-blocks/accounts"
	"github.com/duck-dynasty/duck-blocks/blockchain"
	"github.com/duck-dynasty/duck-blocks/core"
)

func main() {
	fmt.Println("--------------- USER ACCOUNT TESTS ---------------")
	u1 := accounts.NewUserAccount()
	fmt.Println(u1.JSON())
	fmt.Println(core.PublicKeyToHex(u1.PublicKey))
	fmt.Println(core.GetAddress(core.PublicKeyToHex(u1.PublicKey)))
	fmt.Println("")
	fmt.Println("")

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
}
