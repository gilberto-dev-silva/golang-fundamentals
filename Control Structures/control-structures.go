package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 10

	if number == 10 {
		fmt.Println("Number is 10")
	} else if number == 20 {
		fmt.Println("Number is 20")
	} else {
		fmt.Println("Number is not 10 or 20")
	}

	// If init
	if number2 := number; number2 == 10 {
		fmt.Println("Condition is true, Number 2 is 10")
	} else {
		fmt.Println("Condition is false, Number 2 is not 10")
	}
}
