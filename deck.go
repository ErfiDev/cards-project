package main

import "fmt"

type deck [][]byte

func newDeck() deck {
	cards := deck{}

	cardSuite := []string{"Spades", "Diamonds", "Hearts"}
	cardValue := []string{"Ace", "Two", "Three"}

	for i, card := range cardSuite {
		for j, value := range cardValue {
			toByte := []byte(card + " of " + value)
			cards = append(cards, toByte)
			i += j
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
