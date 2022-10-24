package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card  of Four of Cubes, but got %v", d[15])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDesk := newDeckFromFile(fileName)
	if len(loadedDesk) != 16 {
		t.Error("Expected 16 cards in deck, got ", len(loadedDesk))
	}
	os.Remove(fileName)

}
