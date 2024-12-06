package blockchain

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

const difficulty = 5

func TestBlock_calculateHash(t *testing.T) {
	b := &Block{
		previousHash: "previousHash",
		payload:      []byte("payload"),
		timestamp:    1632312000,
		pow:          12345,
	}

	expectedHash := "2509db91908422a87c2de55b0305cafaba20afcc2b050c483e2560a7f3a31fb8"

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

	expectedHash = "696587ed13a8b5f14e7f7b80768f4146dbc0cc66d4366b64f4d55acb4458d695"

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

	expectedHash = "b889a9602ff5b2cc28cba570a94a35a342b609addde78527926daaa93abc6406"

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

// BenchmarkMine is a function for benchmarking the mine method of the Block struct.
func BenchmarkMine(b *testing.B) {
	// Reset the timer to exclude setup time.
	b.ResetTimer()
	// Loop b.N times, which is a way for Go to ask the benchmarking code to keep running until b.StopTimer() is called.
	for i := 0; i < b.N; i++ {
		// Create a new Block instance with predefined values for previousHash, payload, timestamp, and pow.
		block := &Block{
			previousHash: "0",
			payload:      []byte(gofakeit.BitcoinAddress()),
			timestamp:    time.Now().UTC().UnixNano(),
			pow:          0,
		}
		// Perform the mining operation on the block with a predefined difficulty level.
		block.mine(difficulty)
	}
}
