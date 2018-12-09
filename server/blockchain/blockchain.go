package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Hash string
	PrevBlockHash string
	Data string
}

type Blcockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
    b.Hash = hex.EncodeToString(hash[:])
	}
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{
		Data: data,
		PrevBlockHash: prevBlockHash,
	}
	block.setHash()

	return block
}

func (bc *Blcockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks) -1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

func NewBlockchain() *Blcockchain{
	return &Blcockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}