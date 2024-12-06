package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	llpayload = `{
	"from": "0x0000000000000000000000000000000000000000",
	"to": "0x0000000000000000000000000000000000000001",
	"amount": 5
}`
	llpayload2 = `{
	"from": "0x0000000000000000000000000000000000000001",
	"to": "0x0000000000000000000000000000000000000002",
	"amount": 5
}`
	llpayload3 = `{
	"from": "0x0000000000000000000000000000000000000002",
	"to": "0x0000000000000000000000000000000000000003",
	"amount": 5
}`
)

func TestLinkedListBlockchainFunctionality(t *testing.T) {
	chain := CreateLinkedListBlockchain([]byte(llpayload), 4)
	assert.NotNil(t, chain, "Expected a valid blockchain object")

	chain.AddBlock([]byte(llpayload2))
	assert.Equal(t, 2, chain.Size(), "Expected blockchain length to be 2 after adding a block")

	chain.AddBlock([]byte(llpayload3))
	assert.Equal(t, 3, chain.Size(), "Expected blockchain length to be 3 after adding another block")

	assert.True(t, chain.IsValid(), "Expected the blockchain to be valid")
}
