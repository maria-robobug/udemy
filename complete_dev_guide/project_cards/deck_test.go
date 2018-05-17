package main

import (
	"os"
	"testing"
)

const expectedDeckLength = 16

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length == %v, but got %v", expectedDeckLength, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected %v cards in deck, but got %v", expectedDeckLength, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
