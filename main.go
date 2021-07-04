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

	//this is a slice
	Slice := []int{}
	Slice = append(Slice, 10)
	Slice = append(Slice, 15)
	Slice = append(Slice, 20)
	Slice = append(Slice, 25)
	Slice = append(Slice, 30)
	fmt.Println(Slice)

	//this is a array
	Arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(Arr)

	var Arr2 = [2]int{1, 2}
	fmt.Println(Arr2)

	var sliceVar []int
	sliceVar = append(sliceVar, 10)

	for i := 0; i < 10; {
		fmt.Println(i)
	}
}

func PrintDate(year, day, month int) {
	fmt.Printf("%d", year)
	fmt.Printf("%d", month)
	fmt.Printf("%d", day)
}
