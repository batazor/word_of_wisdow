package proofofwork

import (
	"time"

	"github.com/batazor/word_of_wisdom/internal/domain/block"
)

// NewBlock creates and returns Block
func NewBlock(data string, prevHash []byte) (*block.Block, error) {
	block := &block.Block{Timestamp: time.Now().UnixNano(), PrevHash: prevHash, Data: []byte(data)}

	pow, err := NewPoW(block)
	if err != nil {
		return nil, err
	}

	nonce, hash := pow.Work()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block, nil
}
