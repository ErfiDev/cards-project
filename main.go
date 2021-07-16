package main

func main() {
	cards := newDeck()
	cards.print()
	deal(cards, 10)
	
	newTenCar := createRandomCar(10);
	newTenCar.print()


	// newMap := make(map[string]bool);
	// newMap["single?"] = true;
	// fmt.Printf("\n %v \n" , newMap);
}
