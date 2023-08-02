package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of %v, but got %v", 52, len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be %v, but got %v", "Ace of Spades", d[0])
	}
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be %v, but got %v", "King of Diamonds", d[len(d)-1])
	}
}

func SaveDeckAndLoadDeck(t *testing.T) {
	os.Remove("test_nmy_cards")
	d := newDeck()
	saveDeck(d, "test_my_cards")
	cards := loadDeck("test_my_cards")
	if len(cards) != 52 {
		t.Errorf("Expected after load deck length of %v, but got %v", 52, len(d))
	}
}
