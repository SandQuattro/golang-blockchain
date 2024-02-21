package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

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
	generateHash()
	calculateHash() string
	mine(difficulty int)
}

func (b *Block) generateHash() {
	b.hash = b.calculateHash()
}

// calculateHash calculates the hash of the current block.
func (b *Block) calculateHash() string {
	// Concatenate previous hash, payload, timestamp, and proof of work.
	var data []byte
	data = append(data, b.previousHash...)
	data = append(data, b.payload...)
	data = binary.LittleEndian.AppendUint64(data, uint64(b.timestamp))
	data = append(data, strconv.Itoa(b.pow)...)

	// Compute the SHA256 hash of the concatenated data and return it as a hexadecimal string.
	return fmt.Sprintf("%x", sha256.Sum256(data))
}

// mine performs the mining process based on the specified difficulty level.
// If the difficulty is 0, it sets pow to 0 and generates a hash.
// Otherwise, it continues generating a hash until it starts with the required number of zeros.
func (b *Block) mine(difficulty int) {
	// If the difficulty is 0, set pow to 0 and generate a hash
	if difficulty <= 0 {
		b.pow = 0
		b.generateHash()
		return
	}

	// Continue generating a hash until it starts with the required number of zeros
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.generateHash()
	}
}

func (b *Block) encodeData(data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	b.payload = bytes
	return nil
}

func (b *Block) decodeData(result any) error {
	err := json.Unmarshal(b.payload, &result)
	if err != nil {
		return err
	}
	return nil
}
