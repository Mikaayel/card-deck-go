package main

func main() {
	cards := newDeck()
	cards.shuffle()
	// hand, remainingDeck := cards.deal(5)
	// hand.print()
	// remainingDeck.print()
	// saveDeck(cards, "my_cards")
	// cards = loadDeck("my_cards")
	cards.print()
}
