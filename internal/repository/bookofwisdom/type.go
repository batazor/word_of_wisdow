package bookofwisdom

type Repository struct {
	quotes []*Quote
}

// Quote - represents a quote
type Quote struct {
	Author string `json:"Author"`
	Quote  string `json:"Quote"`
}
