package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 56 {
		t.Errorf("expected deck length of 56 but got %d ", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("expected first card to be Ace of Spades but got %s ", cards[0])
	}

	if cards[len(cards)-1] != "Queen of Clubs" {
		t.Errorf("expected second card to be Queen of Club but got %s ", cards[len(cards)-1])
	}

}

func TestDecksFromToFile(t *testing.T) {
	os.Remove("_999testing_file")

	cards := newDeck()
	cards.saveToFile("_999testing_file")

	cardsFromFile := deckFromFile("_999testing_file")

	if len(cardsFromFile) != 56 {
		t.Errorf("expected deck length of 56 but got %d ", len(cardsFromFile))
	}

	os.Remove("_999testing_file")

}
