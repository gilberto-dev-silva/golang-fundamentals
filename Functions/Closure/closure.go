package main

import "fmt"

func closure() func(int) int {
	return func(x int) int {
		return x + 1
	}
}

func main() {
	increment := closure()
	fmt.Println(increment(5))
}
