package main

import "fmt"

func main() {
	var variable1 string = "Variável 1"
	variable2 := "Variável 2"
	fmt.Println("Tipo da variável 1: ", variable1, variable2)

	var (
		variable3 int     = 10
		variable4 float64 = 3.14159
	)
	fmt.Println("Tipo da variável 2: ", variable3, variable4)

	variable5, variable6 := "Variável 5", "Variável 6"
	fmt.Println("Tipo da variável 3: ", variable5, variable6)
}
