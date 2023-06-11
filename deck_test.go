package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
    d := newDeck()
    eDSize := 16
    a := "Ace of Spades"
    c := "Four of Clubs"

    if len(d) != eDSize {
        // t.Errorf("eDSizeected deck length of 16, but got: %d", len(d))
        t.Errorf("eDSizeected deck length of %v, but got: %v", eDSize, len(d))
    }

    if d[0] != a {
        t.Errorf("expected first card in deck to be: %v, got %v", a, d[0])
    }

    if d[len(d) - 1] != c {
        t.Errorf("expected last card in deck to be: %v, got %v", c, d[len(d) - 1])
    }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")
}
