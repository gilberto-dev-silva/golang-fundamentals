package main

import "fmt"

func invertSign(number int) int {
	return number * -1
}

func invertSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	numberInverted := invertSign(number)
	fmt.Println(numberInverted)
	fmt.Println(number)

	newNumber := 40
	fmt.Println(newNumber)
	invertSignalWithPointer(&newNumber)
	fmt.Println(newNumber)

}
