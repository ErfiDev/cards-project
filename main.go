package main

func main() {
	cards := newDeckFromFile("data.txt")
	cards.print()
}
