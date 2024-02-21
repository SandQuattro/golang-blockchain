package blockchain

type Block struct {
	payload []byte
	// previous block hash, will use it for integrity check
	previousHash string
	hash         string
	timestamp    int64
	// proof-of-work computation
	pow int
}

type BlockOperations interface {
	encodeData(data any) error
	decodeData(result any) error
	setHash()
	calculateHash() string
	mine(difficulty int)
}
