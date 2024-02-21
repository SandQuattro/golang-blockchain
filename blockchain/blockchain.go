package blockchain

import "time"

// BLOCKCHAIN

type Blockchain struct {
	// initial block in chain
	initial Block
	chain   []Block
	// difficulty of proof-of-work, number of leading zeros in hash
	difficulty int
}

type Operations interface {
	AddBlock(data []byte)
	IsValid() bool
	Size() int
}

// AddBlock adds a new block to the blockchain with the given data.
func (c *Blockchain) AddBlock(data []byte) {
	// Get the last block to save its hash in the new block
	lastBlock := c.chain[len(c.chain)-1]
	// Create a new block with the provided data and previous hash, and set the timestamp
	newBlock := Block{
		payload:      data,
		previousHash: lastBlock.hash,
		timestamp:    time.Now().UTC().UnixNano(),
	}

	// Apply mining to the new block, computing the proof of work and hash
	// If the difficulty is 0, just compute the hash
	newBlock.mine(c.difficulty)
	// Append the new block to the blockchain
	c.chain = append(c.chain, newBlock)
}

// IsValid checks the validity of the blockchain
func (c *Blockchain) IsValid() bool {
	// Iterate through the chain starting from the second block
	for i := range c.chain[1:] {
		// Get the previous and current blocks
		prevBlock := c.chain[i]
		currentBlock := c.chain[i+1]
		// Check if the previous block's hash matches the current block's previousHash
		// or if the current block's hash matches the hash calculated from its contents
		if prevBlock.hash != currentBlock.previousHash || currentBlock.hash != currentBlock.calculateHash() {
			// If there is a mismatch, the blockchain is not valid
			return false
		}
	}
	// If all blocks are valid, the blockchain is valid
	return true
}

func (c *Blockchain) Size() int {
	return len(c.chain)
}

// CreateBlockchain creates a new blockchain with the given data and difficulty.
//
// It takes a byte slice data and an integer difficulty as parameters and returns a Blockchain.
func CreateBlockchain(data []byte, difficulty int) Blockchain {
	// Create the initial block with the provided data, "0" as the previous hash, and the current timestamp.
	initial := Block{
		payload:      data,
		previousHash: "0",
		timestamp:    time.Now().UTC().UnixNano(),
	}

	// If the difficulty is 0, set the proof of work to 0 and generate the hash for the initial block.
	if difficulty == 0 {
		initial.pow = 0
		initial.generateHash()
	} else {
		// Mine the block and set the hash.
		initial.mine(difficulty)
	}

	// Return a new Blockchain with the initial block as the first block in the chain and the provided difficulty.
	return Blockchain{
		initial:    initial,
		chain:      []Block{initial},
		difficulty: difficulty,
	}
}
