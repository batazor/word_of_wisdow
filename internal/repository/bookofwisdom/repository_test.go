package repository

import (
	"testing"
)

func TestNew(t *testing.T) {
	repo, err := New("data.json")
	if err != nil {
		t.Fatalf("Failed to create repository: %s", err)
	}

	// Test the length
	if len(repo.quotes) != 10 {
		t.Fatalf("Expected 1 quote, got %d", len(repo.quotes))
	}

	// Test the content
	quote := repo.quotes[0]
	if quote.Author != "Albert Einstein" || quote.Quote != "Life is like riding a bicycle. To keep your balance you must keep moving." {
		t.Fatalf("Unexpected quote data: %v", quote)
	}
}
