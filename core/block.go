package core

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"github.com/izqui/functional"
	"github.com/izqui/helpers"
)


type Block struct {
	*Header
	[]*Transactions
  Signature []byte
}

type Header struct {
	PrevBlock  []byte
	MerkelRoot []byte
	Timestamp  uint32
	Nonce      uint32
}

func NewBlock(previousBlock Block) Block {
	header := &Header{PrevBlock: previousBlock.Signature}
	return Block{header, nil, nil}
}

func (b *Block) AddTransaction(t *Transaction) {
	newSlice := b.TransactionSlice.AddTransaction(*t)
	b.TransactionSlice = &newSlice
}

func (b *Block) Sign(keypair *Keypair) []byte {

	s, _ := keypair.Sign(b.Hash())
	return s
}
