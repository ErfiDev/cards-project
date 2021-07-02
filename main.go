package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	now := time.Now()
	PrintDate(now.Year(), now.Day(), int(now.Month()))

	floting := 3.14
	toInt := int(floting)
	fmt.Printf("\n %d \n", toInt)

	randomNum := math.Pi
	fmt.Println(randomNum)

	// Arr := []string;
	// var Arr []string
	// var i int
	// for i = 0; i <= 5; i++ {
	// 	input := ""
	// 	fmt.Scanf("%s", &input)
	// 	Arr = append(Arr, input)
	// }
	// fmt.Println(Arr)
	Arr := []string{"erfan", "ali", "reza"}
	fmt.Println(len(Arr), "length")
}

func PrintDate(year, day, month int) {
	fmt.Printf("%d", year)
	fmt.Printf("%d", month)
	fmt.Printf("%d", day)
}
