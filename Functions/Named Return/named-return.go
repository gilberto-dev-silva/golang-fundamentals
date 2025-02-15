package main

import "fmt"

func calculationsMathematics(x, y int) (sum, subtraction int) {
	sum = x + y
	subtraction = x - y
	return
}

func main() {
	resultSum, resultSubtraction := calculationsMathematics(50, 10)
	fmt.Println(resultSum, resultSubtraction)
}
