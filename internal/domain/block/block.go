package block

import (
	"time"

	"github.com/batazor/word_of_wisdow/internal/domain/proofofwork"
)

type Block struct {
	Timestamp int64
	Hash      []byte
	PrevHash  []byte
	Data      []byte
	Nonce     uint64
}

// NewBlock creates and returns Block
func New(data string, prevHash []byte) (*Block, error) {
	block := &Block{Timestamp: time.Now().UnixNano(), PrevHash: prevHash, Data: []byte(data)}

	pow, err := proofofwork.New(block)
	if err != nil {
		return nil, err
	}

	nonce, hash := pow.Work()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block, nil
}
