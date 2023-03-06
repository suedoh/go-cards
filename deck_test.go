package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T)  {
    d := newDeck()
    firstCard := "Ace of Spades"
    lastCard := "Four of Clubs"

    if len(d) != 16 {
        t.Errorf("Expected deck length of 16 but got %v", len(d))
    }

    if d[0] != firstCard {
        t.Errorf("Expected first card to be %v but got %v", firstCard, d[0])
    }

    if d[len(d) - 1] != lastCard {
        t.Errorf("Expected last card to be %v but got %v", lastCard, d[len(d) - 1])
    }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T)()  {
    testFile := "_decktesting"

    os.Remove(testFile)

    deck := newDeck()
    deck.saveToFile(testFile)

    loadedDeck := newDeckFromFile(testFile)

    if len(loadedDeck) != 16 {
        t.Errorf("Expected 16 cards but received %v", len(loadedDeck))
    }

    os.Remove(testFile)
}
