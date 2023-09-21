package proofofwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	data := "some data"
	prevHash := []byte{}

	newBlock, err := NewBlock(data, prevHash)
	assert.NoError(t, err)

	assert.NotEmpty(t, newBlock.Hash)
	assert.Equal(t, prevHash, newBlock.PrevHash)
	assert.Equal(t, []byte(data), newBlock.Data)

	// Validate proof-of-work
	pow, err := NewPoW(newBlock)
	assert.NoError(t, err)
	assert.True(t, pow.Verify())
}
