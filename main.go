package main

func main() {
	cards := newDeck()

	hand, remainningDeck := deal(cards, 2)

	hand.print()
	remainningDeck.print()
}
