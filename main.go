package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

type Blockchain struct {
	Block []*Block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	header := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(header)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Timestamp: 0, Data: []byte(data), PrevBlockHash: prevBlockHash, Hash: []byte{}}
	block.SetHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Block[len(bc.Block)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Block = append(bc.Block, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{Block: []*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Block {
		println("Prev. hash: ")
		println(block.PrevBlockHash)
		println("Data: ")
		println(block.Data)
		println("Hash: ")
		println(block.Hash)
	}
}
