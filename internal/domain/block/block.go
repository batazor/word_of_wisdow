package block

type Block struct {
	Timestamp int64
	Hash      []byte
	PrevHash  []byte
	Data      []byte
	Nonce     uint64
}
