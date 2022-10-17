package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Block struct {
	transaction string
	index int
	nonce int
	hash string
	previousHash string
	timestamp time.Time
}

type Blockchain struct {
	genesisBlock Block
	chain []Block
}

func calculateHash(block Block) string {
	hash := block.transaction + string(block.nonce) + string(block.index) + block.hash + block.previousHash + block.timestamp.String()
	blockHash := sha256.Sum256([]byte(hash))
	return fmt.Sprintf("%x", blockHash)
}

func createBlockchain() Blockchain {
	genesisBlock := Block{
		transaction: "none",
		index: 0,
		timestamp: time.Now(),
		nonce: rand.Intn(100),
		previousHash: "0",
	}
	genesisBlock.hash = calculateHash(genesisBlock)

	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
	}
}

func addBlock(transaction string, blockchain *Blockchain) {
	oldBlock := blockchain.chain[len(blockchain.chain)-1]

	newBlock := Block{
		transaction: transaction,
		index: oldBlock.index + 1,
		timestamp: time.Now(),
		nonce: rand.Intn(100),
		previousHash: oldBlock.hash,
	}

	newBlock.hash = calculateHash(newBlock)

	blockchain.chain = append(blockchain.chain, newBlock)
}

func verifyChain(blockChain *Blockchain) bool{
	for i := 0 ; i < len(blockChain.chain)-1 ; i++ {
		previousBlock := blockChain.chain[i]
		currentBlock := blockChain.chain[i+1]
		if currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func listBlock(blockChain *Blockchain) {
	for i:=0 ; i<len(blockChain.chain) ; i++ {
		fmt.Printf("Transaction: %v\n", blockChain.chain[i].transaction)
		fmt.Printf("Index: %v\n", blockChain.chain[i].index)
		fmt.Printf("Nonce: %v\n", blockChain.chain[i].nonce)
		fmt.Printf("Timestamp: %v\n", blockChain.chain[i].timestamp)
		fmt.Printf("Previous Hash: %v\n", blockChain.chain[i].previousHash)
		fmt.Printf("Current Block Hash: %v\n\n", blockChain.chain[i].hash)
	}
}

func changeBlock(transactionChange string, blockChain *Blockchain, index int) {
	blockChain.chain[index].transaction = transactionChange
}

func Menu(){
	var blockchain Blockchain
	in := bufio.NewReader(os.Stdin)

	for {
		var choice int
		fmt.Print("******##### Hitcoin Portal ##### *******\n1) Create Blockchain\n2) Add a new Block\n3) List all blocks\n4) Change transaction in a block\n5) Verify Blockchain\n6) Exit\n")
		fmt.Print("Enter Choice Here: ")
		fmt.Scanf("%d\n" , &choice)

		switch choice {
			case 1:
				blockChain := createBlockchain()
				blockchain = blockChain
				fmt.Print("Your Blockchain has been created\n")
			case 2:
				fmt.Print("Input Transaction: ")
				input, _ := in.ReadString('\n')
				addBlock(input, &blockchain)
				fmt.Print("Your block has been added\n")
			case 3:
				listBlock(&blockchain)
			case 4:
				var index int
				fmt.Print("Input Transaction: ")
				input, _ := in.ReadString('\n')
				fmt.Print("Input Block Index: ")
				fmt.Scanf("%d\n", &index)
				changeBlock(input, &blockchain, index)
				fmt.Print("Block Changed Successfully\n")
			case 5:
				verify := verifyChain(&blockchain)
				fmt.Print("Chain Authentic: %v\n", verify)
			case 6:
				os.Exit(0)	
		}
	}
}

func main(){
	Menu()
}
