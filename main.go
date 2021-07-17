package main

import (
	"fmt"
	"time"
)

func main() {
	cards := newDeck()
	cards.print()

	newCars := createRandomCar(3)
	newCars.print()

	day := time.Now().Day()
	fmt.Println(day)

}
