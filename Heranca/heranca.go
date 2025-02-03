package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course  string
	faculty string
}

func main() {
	fmt.Println("Herança")

	p1 := person{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := student{p1, "Engenharia", "USP"}
	fmt.Println(e1.height)
}
