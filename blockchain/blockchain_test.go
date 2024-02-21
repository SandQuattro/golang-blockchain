package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	payload = `{
	"from": "0x0000000000000000000000000000000000000000",
	"to": "0x0000000000000000000000000000000000000001",
	"amount": 5
}`
	payload2 = `{
	"from": "0x0000000000000000000000000000000000000001",
	"to": "0x0000000000000000000000000000000000000002",
	"amount": 5
}`
	payload3 = `{
	"from": "0x0000000000000000000000000000000000000002",
	"to": "0x0000000000000000000000000000000000000003",
	"amount": 5
}`
)

func TestBlockchainFunctionality(t *testing.T) {

	chain := CreateBlockchain([]byte(payload), 4)
	assert.NotNil(t, chain, "Expected a valid blockchain object")

	chain.AddBlock([]byte(payload2))
	assert.Equal(t, 2, chain.Size(), "Expected blockchain length to be 2 after adding a block")

	chain.AddBlock([]byte(payload3))
	assert.Equal(t, 3, chain.Size(), "Expected blockchain length to be 3 after adding another block")

	assert.True(t, chain.IsValid(), "Expected the blockchain to be valid")
}
