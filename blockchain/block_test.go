package blockchain

import (
	"math"
	"strings"
	"testing"
)

func TestBlock_calculateHash(t *testing.T) {
	b := &Block{
		previousHash: "previousHash",
		payload:      []byte("payload"),
		timestamp:    1632312000,
		pow:          12345,
	}

	expectedHash := "b6817dbb1fd542b05db571d310fea0484fed286995ff47608d7c1d9b833e81dc"

	actualHash := b.calculateHash()

	if actualHash != expectedHash {
		t.Errorf("Expected hash: %s, but got: %s", expectedHash, actualHash)
	}

	// Testing the hash calculation with an empty previous hash and payload
	b = &Block{
		previousHash: "",
		payload:      []byte(""),
		timestamp:    0,
		pow:          0,
	}

	expectedHash = "d26493520dce67d4a71c93ab2827cb8dce6a25353f6c6d08d9bd5585a549e628"

	actualHash = b.calculateHash()

	if actualHash != expectedHash {
		t.Errorf("Expected hash: %s, but got: %s", expectedHash, actualHash)
	}

	// Testing the hash calculation with a large timestamp and proof of work
	b = &Block{
		previousHash: "previousHash",
		payload:      []byte("payload"),
		timestamp:    math.MaxInt64, // max uint64 value
		pow:          2147483647,    // max int32 value
	}

	expectedHash = "0dfa88eaeb485c965deda4442182ddddb118cd837ef008604b734e3e556797ea"

	actualHash = b.calculateHash()

	if actualHash != expectedHash {
		t.Errorf("Expected hash: %s, but got: %s", expectedHash, actualHash)
	}
}

func TestMine(t *testing.T) {
	block := &Block{}

	// Test mining with difficulty 0
	block.mine(0)
	if block.pow != 0 {
		t.Errorf("Expected pow to be 0, got %d", block.pow)
	}

	// Test mining with difficulty 1
	block.mine(1)
	if !strings.HasPrefix(block.hash, "0") {
		t.Errorf("Expected hash to start with 0, got %s", block.hash)
	}

	// Test mining with difficulty 2
	block.mine(2)
	if !strings.HasPrefix(block.hash, "00") {
		t.Errorf("Expected hash to start with 00, got %s", block.hash)
	}

	// Test mining with negative difficulty
	block.mine(-1)
	// No need to check for specific result, just ensuring that it doesn't cause any errors
}
