package main

import (
	"fmt"
)

func main() {
	humans := human{}
	humans = humans.addNew("erfan", "hanifezade", "developer", 17)
	humans = humans.addNew("maziar", "rezaee", "mechanic", 22)
	fmt.Println(humans)
}
