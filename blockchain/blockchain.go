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

func (c *Blockchain) AddBlock(data []byte) {
	// получаем последний блок для сохранения предыдущего hash в новом блоке
	lastBlock := c.chain[len(c.chain)-1]
	newBlock := Block{
		payload:      data,
		previousHash: lastBlock.hash,
		timestamp:    time.Now().UTC().UnixNano(),
	}

	// применяем майнинг к новому блоку, вычисляется pow и hash
	// при нулевой сложности просто вычисляем hash
	newBlock.mine(c.difficulty)
	c.chain = append(c.chain, newBlock)
}

func (c *Blockchain) IsValid() bool {
	for i := range c.chain[1:] {
		prevBlock := c.chain[i]
		currentBlock := c.chain[i+1]
		if prevBlock.hash != currentBlock.previousHash || currentBlock.hash != currentBlock.calculateHash() {
			return false
		}
	}
	return true
}

func (c *Blockchain) Size() int {
	return len(c.chain)
}

func CreateBlockchain(data []byte, difficulty int) Blockchain {
	initial := Block{
		payload:      data,
		previousHash: "0",
		timestamp:    time.Now().UTC().UnixNano(),
	}
	if difficulty == 0 {
		initial.pow = 0
		initial.generateHash()
	} else {
		// майним блок, устанавливаем hash
		initial.mine(difficulty)
	}

	return Blockchain{
		initial:    initial,
		chain:      []Block{initial},
		difficulty: difficulty,
	}
}
