package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

type block struct {
	Hash     string
	Data     string
	Noonce   int
	PrevHash string
	Genesis  bool
}

type Blockchain struct {
	B_chain []*block
}

func NewBlock(transaction string, nonce int, previousHash string) *block {
	block_obj := &block{
		CalculateHash(transaction),
		transaction,
		nonce,
		previousHash,
		false,
	}
	return block_obj
}

func (b_chain *Blockchain) AddBlock(transaction string) {
	if len(b_chain.B_chain) == 0 {
		new_data := NewBlock(transaction, rand.Intn(1000), "")
		new_data.Genesis = true
		b_chain.B_chain = append(b_chain.B_chain, new_data)
	} else {
		new_data := NewBlock(transaction, rand.Intn(1000), b_chain.B_chain[len(b_chain.B_chain)-1].Hash)
		b_chain.B_chain = append(b_chain.B_chain, new_data)
	}

}

func (b_chain *Blockchain) DisplayBlocks() {
	for i := range b_chain.B_chain {
		fmt.Println("==============================================")
		fmt.Println("Current Block Hash : ", b_chain.B_chain[i].Hash)
		if b_chain.B_chain[i].Genesis { // If starting block
			fmt.Println("This is the starting block thus no previous hashes")
		} else {
			fmt.Println("Previous Block Hash : ", b_chain.B_chain[i].PrevHash)
		}
		fmt.Println("Noonce value : ", b_chain.B_chain[i].Noonce)
		fmt.Println("Current Block Data : ", b_chain.B_chain[i].Data)
		fmt.Println("==============================================")
	}
}

func (b_chain *Blockchain) ChangeBlock(index int, transaction string) {
	b_chain.B_chain[index].Data = transaction
	b_chain.B_chain[index].Hash = CalculateHash(transaction)
	b_chain.VerifyChain()
}

func (b_chain *Blockchain) VerifyChain() {
	fmt.Println("Verifying Chain")
	// Goes through the entire chain and re-calculates hashes if changes are made
	for i := range b_chain.B_chain {
		if i != 0 {
			if b_chain.B_chain[i].PrevHash != b_chain.B_chain[i-1].Hash {
				b_chain.B_chain[i].PrevHash = b_chain.B_chain[i-1].Hash
			}
		}
		if b_chain.B_chain[i].Hash != CalculateHash(b_chain.B_chain[i].Data) {
			b_chain.B_chain[i].Hash = CalculateHash(b_chain.B_chain[i].Data)
		}
	}
}

func CalculateHash(stringToHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}
