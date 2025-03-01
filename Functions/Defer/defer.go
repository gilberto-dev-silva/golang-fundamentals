package main

import "fmt"

func functionOne() {
	fmt.Println("Function One")
}

func functionTwo() {
	fmt.Println("Function Two")
}

func studentAproved(n1, n2 float32) bool {
	defer fmt.Printf("Student Aproved: %v, %v\n", n1, n2)
	average := n1 + n2/2

	if average >= 7 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Student Aproved:", studentAproved(8, 9))
}
