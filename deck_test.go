package main

import (
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
