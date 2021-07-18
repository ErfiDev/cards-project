package main

import (
	"fmt"
	"math/rand"
)

type carStruct struct {
	name  string
	color string
}
type car []carStruct

func createCar(name, color string) car {
	newModel := carStruct{
		name:  name,
		color: color,
	}

	newCar := car{newModel}
	return newCar
}

func createRandomCar(count int) car {
	names := []string{"opel", "bmw", "dodge", "mustang", "honda"}
	colors := []string{"blue", "white", "red", "orange", "yellow"}

	newCar := car{}

	for i := 0; i <= count; i++ {
		randomForName := randomNum(5)
		randomForColor := randomNum(5)
		newCar = append(newCar, carStruct{
			name:  names[randomForName],
			color: colors[randomForColor],
		})
	}

	return newCar
}

func (c car) print() {
	fmt.Println(c)
}

func randomNum(index int) int {
	return rand.Intn(index)
}
