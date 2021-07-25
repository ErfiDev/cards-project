package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuite := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

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

func (d deck) toByteSlice() []byte {
	toString := strings.Join(d, ",")

	byteSlice := []byte(toString)
	return byteSlice
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

func newDeckFromFile(filename string) deck {
	slice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	toString := string(slice)
	toStrSlice := strings.Split(toString, ",")
	return deck(toStrSlice)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		randomInt := r.Intn(len(d) - 1)
		d[index], d[randomInt] = d[randomInt], d[index]
	}
}
