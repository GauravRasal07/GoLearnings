package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("The length of deck must be 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("The first card is not Ace of Spades, It is: %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("The last card is not Four of Clubs, It is: %v", d[len(d) - 1])
	}
}

func TestWriteToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	deck := newDeck()
	deck.writeToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")

	if (len(loadedDeck) != 16) {
		t.Errorf("The length of deck must ne 16 but got %v", len(loadedDeck))
	}
	
	os.Remove("_decktesting.txt")
}