package main

func main() {
	cardsFromFile := newDeckFromFile("data.txt")
	cardsFromFile.print()
}
