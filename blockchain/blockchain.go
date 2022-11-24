package blockchain

import (
	"time"

	"github.com/duck-dynasty/duck-blocks/accounts"
	"github.com/duck-dynasty/duck-blocks/transactions"
)

type Block struct {
	Index         int                              `json:"index"`
	Proof         string                           `json:"proof"`
	Timestamp     int64                            `json:"timestamp"`
	Transactions  []transactions.DuckatTransaction `json:"transactions"`
	PreviousHash  string                           `json:"previous_hash"`
	AccountStates map[string]accounts.AccountState `json:"account_states"`
}

type Blockchain struct {
	Chain               []Block                          `json:"chain"`
	CurrentTransactions []transactions.DuckatTransaction `json:"current_transactions"`
	CurrentStates       map[string]accounts.AccountState `json:"account_states"`
}

func (b *Blockchain) LastBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

func (b *Blockchain) UpdateFromState(address string, value int64) {
	state, _ := b.CurrentStates[address]
	state.Nonce += 1
	state.Value -= value
	b.CurrentStates[address] = state
}

func (b *Blockchain) UpdateToState(address string, value int64) {
	state, ok := b.CurrentStates[address]
	if ok {
		state.Value += value
	} else {
		state = accounts.AccountState{Address: address, Value: value, Code: nil, Data: nil, Nonce: 0}
	}

	b.CurrentStates[address] = state
}

func (b *Blockchain) NewTransaction(newTransaction transactions.DuckatTransaction) bool {
	accountState, ok := b.CurrentStates[newTransaction.Params.From]

	if newTransaction.VerifyTransaction() {
		if ok {
			if accountState.Value >= newTransaction.Params.Amount {
				b.CurrentTransactions = append(b.CurrentTransactions, newTransaction)
				b.UpdateFromState(newTransaction.Params.From, newTransaction.Params.Amount)
				b.UpdateToState(newTransaction.Params.To, newTransaction.Params.Amount)
			}
		}
	}

	return false
}

func (b *Blockchain) NewBlock(proof string, previousHash string) Block {
	block := Block{
		Index:         len(b.Chain) + 1,
		Proof:         proof,
		Timestamp:     time.Now().Unix(),
		Transactions:  b.CurrentTransactions,
		PreviousHash:  previousHash,
		AccountStates: b.CurrentStates,
	}

	b.CurrentTransactions = nil
	b.Chain = append(b.Chain, block)

	return block
}

func NewBlockchain() Blockchain {
	blockchain := Blockchain{
		Chain:               []Block{},
		CurrentTransactions: []transactions.DuckatTransaction{},
		CurrentStates:       make(map[string]accounts.AccountState),
	}

	_ = blockchain.NewBlock("100", "1")

	return blockchain
}
