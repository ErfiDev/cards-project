package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	cards := newDeck()
	cards.print()

	newCars := createRandomCar(3)
	newCars.print()

	day := time.Now().Day()
	fmt.Printf("%d day \n", day)

	humans := human{}
	humans = humans.addNew("erfan", "hanifezade", "developer", 17)
	humans = humans.addNew("maziar", "rezaee", "mechanic", 22)
	fmt.Println(humans[1].family)

	result, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error")
	}

	os.Stdout.Write(result)
}
