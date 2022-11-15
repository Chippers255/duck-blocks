package core

import (
	"bytes"
	"crypto/rsa"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

// Converts an RSA Public Key to a hex string for easier usage. Struct serialization found in https://stackoverflow.com/questions/28020070/golang-serialize-and-deserialize-back
func PublicKeyToHex(pk rsa.PublicKey) string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(pk)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}
	return hex.EncodeToString(b.Bytes())
}

// Converts a hex string to a RSA Public Key for signature verification. Struct deserialization found in https://stackoverflow.com/questions/28020070/golang-serialize-and-deserialize-back
func HexToPublicKey(hexString string) rsa.PublicKey {
	m := rsa.PublicKey{}
	by, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println(`failed base64 Decode`, err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
	return m
}

func GetAddress(publicKeyHex string) string {
	last40 := publicKeyHex[len(publicKeyHex)-40:]
	address := "0x" + string(last40)
	return address
}
