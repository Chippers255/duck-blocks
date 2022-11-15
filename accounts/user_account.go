package accounts

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"os"

	core "github.com/duck-dynasty/duck-blocks/core"
	transactions "github.com/duck-dynasty/duck-blocks/transactions"
)

type UserAccount struct {
	State      AccountState   `json:"state"`
	PrivateKey rsa.PrivateKey `json:"private_key"`
	PublicKey  rsa.PublicKey  `json:"public_key"`
	Address    string         `json:"address"`
}

func NewUserAccount() *UserAccount {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey
	address := core.GetAddress(core.PublicKeyToHex(publicKey))

	state := AccountState{Address: address, Value: 0, Code: nil, Data: nil, Nonce: 0}
	user := UserAccount{Address: address, PrivateKey: *privateKey, PublicKey: publicKey, State: state}

	return &user
}

func SaveKeyFile(fileName string, keyBlock *pem.Block) {
	pemKeyFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = pem.Encode(pemKeyFile, keyBlock)
	if err != nil {
		panic(err)
	}

	pemKeyFile.Close()
}

func (u UserAccount) SaveKeys() {
	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&u.PrivateKey),
	}
	SaveKeyFile("private_key.pem", pemPrivateBlock)

	var pemPublicBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&u.PublicKey),
	}
	SaveKeyFile("public_key.pem", pemPublicBlock)
}

// A method that returns a minimal JSON string of the user account.
func (u UserAccount) JSON() string {
	value, _ := json.Marshal(u)
	return string(value)
}

func (u UserAccount) CreateTransaction(to string, amount int64) transactions.DuckatTransaction {
	publicKey := core.PublicKeyToHex(u.PublicKey)
	p := transactions.DuckatParams{Amount: amount, From: u.Address, MaxFee: 12345, Nonce: u.State.Nonce, To: to, Type: 127, Key: publicKey}
	t := transactions.DuckatTransaction{Params: &p}
	transactions.SignTransaction(&t, &u.PrivateKey)
	return t
}
