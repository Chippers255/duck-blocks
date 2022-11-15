package transactions

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

func TransactionToBytes(transactionParams *DuckatParams) []byte {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(transactionParams)
	if err != nil {
		fmt.Println(`failed encode transaction instance`, err)
	}
	return buffer.Bytes()
}

func HashTransaction(transactionParams *DuckatParams) []byte {
	params := TransactionToBytes(transactionParams)

	paramHash := sha256.New()
	_, err := paramHash.Write(params)
	if err != nil {
		panic(err)
	}

	return paramHash.Sum(nil)
}

func SignTransaction(transactionInstance *DuckatTransaction, PrivateKey *rsa.PrivateKey) {
	paramHashSum := HashTransaction(transactionInstance.Params)

	signature, err := rsa.SignPSS(rand.Reader, PrivateKey, crypto.SHA256, paramHashSum, nil)
	if err != nil {
		panic(err)
	}

	timestamp := time.Now().Format(time.RFC3339)
	signature_string := signature

	transactionInstance.Signature = &DuckatSignature{Value: signature_string, Timestamp: timestamp}
}
