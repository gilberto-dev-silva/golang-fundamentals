package main

import "fmt"

func main() {
	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println("Variable 1:", variable1)
	fmt.Println("Variable 2:", variable2)
	variable1++
	fmt.Println("Variable 1 alterada:", variable1)

	// Pointers is referenced memory address
	var variable3 int = 100
	var variable4 *int = &variable3

	fmt.Println("Variable 3:", variable3)
	fmt.Println("Variable 4:", variable4)  // Pointer value
	fmt.Println("Variable 4:", *variable4) // Dereferencing a pointer
}
