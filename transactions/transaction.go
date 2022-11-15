package transactions

import (
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"fmt"

	core "github.com/duck-dynasty/duck-blocks/core"
)

type Transaction interface {
	GetJSON() string
	VerifyTransaction() string
}

type DuckatParams struct {
	Amount int64  `json:"amount"`  // The amount of Duckats being sent in Duckatons
	From   string `json:"from"`    // The public address of the transaction sender
	MaxFee int64  `json:"max_fee"` // The maximum amount of bread fees a user is willing to spend on a transaction
	Nonce  int64  `json:"nonce"`   // The transaction number of the sender
	To     string `json:"to"`      // The public address of the transaction receiver
	Type   int8   `json:"type"`    // The transaction type, always 127 in this case
	Key    string `json:"key"`     // The hex representation of the senders public key
}

type DuckatSignature struct {
	Value     []byte `json:"value,omitempty"`     // sdfsdf
	Timestamp string `json:"timestamp,omitempty"` // dfsfd
}

type DuckatTransaction struct {
	Params    *DuckatParams    `json:"params,omitempty"`    // sfdsa
	Signature *DuckatSignature `json:"signature,omitempty"` // asdsf
}

// A method that returns a minimal JSON string of the Duckat Transaction.
func (t DuckatTransaction) GetJSON() string {
	value, _ := json.Marshal(t)
	return string(value)
}

func (t DuckatTransaction) VerifyTransaction() bool {
	paramHashSum := HashTransaction(t.Params)
	publicKey := core.HexToPublicKey(t.Params.Key)

	err := rsa.VerifyPSS(&publicKey, crypto.SHA256, paramHashSum, t.Signature.Value, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return false
	}
	return true
}
