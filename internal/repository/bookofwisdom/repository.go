package bookofwisdom

import (
	"encoding/json"
	"os"
)

// New - creates and returns new Repository
func New() (*Repository, error) {
	// read quotes from file
	raw, err := os.ReadFile("data.json")
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
