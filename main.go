package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	byteSlice := cards.toByteSlice()
	fmt.Println(byteSlice)
}
