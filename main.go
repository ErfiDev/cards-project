package main

import (
	"fmt"
)

func main() {
	condition := true
	x := 0
	for condition {
		fmt.Println(x)
		x++
		if x == 10 {
			fmt.Println("break")
			break
		}
	}
}
