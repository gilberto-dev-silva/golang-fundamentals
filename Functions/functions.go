package main

import "fmt"

func sum(a int8, b int8) int8 {
	return a + b
}

func calculationMathematical(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}

func main() {
	result := sum(10, 10)
	fmt.Println(result)

	var fn = func() {
		fmt.Println("Function type")
	}

	fn()

	var fnText = func(txt string) {
		fmt.Println(txt)
	}

	fnText("Function with return")

	resultSum, resultSubtraction := calculationMathematical(10, 5)
	fmt.Println(resultSum, resultSubtraction)
}
