package transactions

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Transaction interface {
	GetJSON() string
	SignTransaction() string
}

type DuckartParams struct {
	Amount int64  `json:"amount"`  // The amount of Duckats being sent in Duckatons
	From   string `json:"from"`    // The public address of the transaction sender
	MaxFee int64  `json:"max_fee"` // The maximum amount of bread fees a user is willing to spend on a transaction
	Nonce  int64  `json:"nonce"`   // The transaction number of the sender
	To     string `json:"to"`      // The public address of the transaction receiver
	Type   int8   `json:"type"`    // The transaction type, always 127 in this case
}

type DuckatSignature struct {
	Value     []byte `json:"value,omitempty"`     // sdfsdf
	Timestamp string `json:"timestamp,omitempty"` // dfsfd
}

type DuckatTransaction struct {
	Params    *DuckartParams   `json:"params,omitempty"`    // sfdsa
	Signature *DuckatSignature `json:"signature,omitempty"` // asdsf
}

func (t *DuckatTransaction) SignTransaction(PrivateKey *rsa.PrivateKey) {
	msg := []byte(t.GetJSON())

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, PrivateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	timestamp := time.Now().Format(time.RFC3339)
	signature_string := signature

	t.Signature = &DuckatSignature{Value: signature_string, Timestamp: timestamp}
	fmt.Println("finished signing")
}

// A method that returns a minimal JSON string of the Duckat Transaction.
func (t DuckatTransaction) GetJSON() string {
	value, _ := json.Marshal(t)
	return string(value)
}
