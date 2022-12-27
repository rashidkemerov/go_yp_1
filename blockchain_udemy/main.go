package main

import (
	"fmt"
	"log"
	"time"
)

// block

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}
func (b *Block) Print() {
	fmt.Printf("timestamp: 	 %d\n", b.timestamp)
	fmt.Printf("nonce: 		 %d\n", b.nonce)
	fmt.Printf("previousHash: %s\n", b.previousHash)
	fmt.Printf("transactions: %s\n", b.transactions)
}

// blockchain

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	//log.Println("test")
	b := NewBlock(0, "init hash")
	b.Print()
	blockChain := NewBlockchain()
	fmt.Println(blockChain)
}
