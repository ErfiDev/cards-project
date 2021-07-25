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
}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	deck.shuffle()

	if len(deck) != 16 {
		t.Errorf("Expected shuffled deck length of 16, but got %v", len(deck))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktest")

	deck := newDeck()
	deck.saveToFile("_decktest")

	readTestFile := newDeckFromFile("_decktest")
	if len(readTestFile) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(readTestFile))
	}

	os.Remove("_decktest")
}
