package main

import "fmt"

type user struct {
	name        string
	age         int8
	addressUser address
}

type address struct {
	street string
	zip    int8
}

func main() {

	ad := address{
		street: "Rua dos Laranjais",
		zip:    98,
	}

	u1 := user{
		name:        "Gilberto",
		age:         20,
		addressUser: ad,
	}

	fmt.Println(u1)

	var u2 user
	u2.name = "Gilberto Silva"
	u2.age = 70

	fmt.Println(u2)
}
