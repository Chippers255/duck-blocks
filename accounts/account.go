package accounts

import transactions "github.com/duck-dynasty/duck-blocks/transactions"

type Account interface {
	JSON() string
	SaveKeys()
	CreateTransaction() transactions.DuckatTransaction
}

type AccountState struct {
	Address string `json:"address"` // The public address of the account, for users this will be the public key
	Value   int64  `json:"value"`   // The current value of the account in duckatons
	Code    []byte `json:"code"`    // Empty for users, the compiled code for contracts
	Data    []byte `json:"data"`    // Constant and variable value states for the account, empty for users
	Nonce   int64  `json:"nonce"`   // The number of transactions sent from this account
}
