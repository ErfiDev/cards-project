package main

type humanSchema struct {
	name   string
	age    int
	family string
	job    string
}
type human []humanSchema

func (h human) addNew(name, family, job string, age int) human {
	h = append(h, humanSchema{
		name,
		age,
		family,
		job,
	})

	return h
}
