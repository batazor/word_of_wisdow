package repository

import (
	"encoding/json"
	rand2 "math/rand"
	"os"
)

// New - creates and returns new Repository
func New(uri string) (*Repository, error) {
	// read quotes from file
	raw, err := os.ReadFile(uri)
	if err != nil {
		return nil, err
	}

	// parse quotes
	data := []*Quote{}
	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	return &Repository{
		quotes: data,
	}, nil
}

// List - return list of quotes
func (r *Repository) List() ([]*Quote, error) {
	return r.quotes, nil
}

// GetRandomItem - return random quote
func (r *Repository) GetRandomItem() (*Quote, error) {
	rand := rand2.Intn(len(r.quotes))
	return r.quotes[rand], nil
}
