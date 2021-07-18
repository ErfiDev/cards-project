package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuite := []string{"Spades", "Diamonds", "Hearts"}
	cardValue := []string{"Ace", "Two", "Three"}

	for i, card := range cardSuite {
		for j, value := range cardValue {
			cards = append(cards, card+" of "+value)
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
