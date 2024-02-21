package blockchain

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
