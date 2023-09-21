package proofofwork

type PoW interface {
	Work() (uint64, []byte)
	Verify() bool
}
