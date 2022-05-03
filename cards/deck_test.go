package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck to start with Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected deck to end with King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"
	os.Remove(testFileName)

	deck := newDeck()

	deck.saveToFile(testFileName)

	deckFromFile := newDeckFromFile(testFileName)

	if len(deckFromFile) != 52 {
		t.Errorf("Expected deck from file length to be 52, but got %v", len(deckFromFile))
	}

	os.Remove(testFileName)
}
