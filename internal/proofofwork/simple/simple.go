package simple

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"

	"github.com/batazor/word_of_wisdom/internal/domain/block"
)

type PoW struct {
	block  *block.Block
	target *big.Int

	// targetBits defines how many leading zeros we need.
	targetBits uint64
	maxNonce   uint64
}

// New returns a new PoW instance.
func New(b *block.Block) (*PoW, error) {
	pow := &PoW{
		block:      b,
		target:     big.NewInt(1),
		targetBits: 24,
		maxNonce:   math.MaxInt64,
	}

	// Lsh sets z = x << n and returns z.
	pow.target.Lsh(pow.target, uint(256-pow.targetBits))

	return pow, nil
}

// prepareData combines block fields into a bytes array.
func (p PoW) prepareData(nonce uint64) []byte {
	data := bytes.Join(
		[][]byte{
			p.block.PrevHash,
			p.block.Data,
			IntToHex(p.block.Timestamp),
			IntToHex(int64(p.targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (p PoW) Work() (uint64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := uint64(0)

	for nonce < p.maxNonce {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (p PoW) Verify() bool {
	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)

	b := big.Int{}
	hashInt := b.SetBytes(hash[:])

	return hashInt.Cmp(p.target) == -1
}
