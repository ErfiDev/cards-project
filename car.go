package main

import (
	"fmt"
	"math/rand"
)

type carStruct struct {
	name string
	color string
} 
type car []carStruct

func createCar(name , color string) car {
	newModel := carStruct{
		name: name,
		color: color,
	}

	newCar := car{newModel};
	return newCar;
}

func createRandomCar(count int) car {
	names := []string{"opel","bmw","dodge","mustang","honda"};
	colors := []string{"blue","white","red","orange","yellow"};

	newCar := car{};

	for i := 0 ; i <= count ; i++ {
		random := rand.Intn(5);
		newCar = append(newCar, carStruct{
			name : names[random],
			color: colors[random],
		})	
	}

	return newCar
}

func (c car) print() {
	fmt.Println(c);
}