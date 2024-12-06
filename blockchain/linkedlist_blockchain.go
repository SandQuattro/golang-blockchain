package blockchain

import (
	"container/list"
	"time"
)

type LinkedListBlockchain struct {
	chain      *list.List
	difficulty int
}

func CreateLinkedListBlockchain(data []byte, difficulty int) *LinkedListBlockchain {
	initialBlock := Block{
		payload:      data,
		previousHash: "0",
		timestamp:    time.Now().UTC().UnixNano(),
	}

	if difficulty == 0 {
		initialBlock.pow = 0
		initialBlock.generateHash()
	} else {
		initialBlock.mine(difficulty)
	}

	blockchain := &LinkedListBlockchain{
		difficulty: difficulty,
		chain:      list.New(),
	}

	blockchain.chain.PushBack(&initialBlock)

	return blockchain
}

func (c *LinkedListBlockchain) AddBlock(data []byte) {
	lastBlock := c.chain.Back().Value.(*Block)
	newblock := Block{
		payload:      data,
		previousHash: lastBlock.hash,
		timestamp:    time.Now().UTC().UnixNano(),
	}

	newblock.mine(c.difficulty)

	c.chain.PushBack(&newblock)
}

func (c *LinkedListBlockchain) Size() int {
	return c.chain.Len()
}

func (c *LinkedListBlockchain) IsValid() bool {
	first := c.chain.Front()
	if first == nil {
		return false
	}
	el := first.Next()

	for ; el != nil; el = el.Next() {
		prevBlock := el.Prev().Value.(*Block)
		currentBlock := el.Value.(*Block)
		if prevBlock.hash != currentBlock.previousHash && currentBlock.hash != currentBlock.calculateHash() {
			return false
		}
	}

	// If all blocks are valid, the blockchain is valid
	return true
}
