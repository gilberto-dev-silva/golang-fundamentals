package main

import "fmt"

func main() {
	// Arithmetic Operators
	sum := 1 + 2
	subtraction := 1 - 2
	division := 1 / 2
	multiplication := 1 * 2
	modulus := 1 % 2

	fmt.Println(sum, subtraction, division, multiplication, modulus)

	var num1 int8 = 10
	var num2 int8 = 20
	result := num1 + num2

	fmt.Println(result)
	// Finally arithmetic

	// Operations Attribution
	var variable1 string = "Assigning value one"
	variable2 := "Assigning value two"
	fmt.Println(variable1)
	fmt.Println(variable2)
	// Finally Operations Attribution

	// Operation Relational
	fmt.Println(1 == 1) // true
	fmt.Println(1 != 1) // false
	fmt.Println(1 > 1)  // false
	fmt.Println(1 < 1)  // false
	fmt.Println(1 >= 1) // true
	fmt.Println(1 <= 1) // true
	// Finally Operation Relational

	// Operation Logical
	fmt.Println("--------------------------------")
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
	// Finally Operation Logical

	// Operation Unitary
	fmt.Println("--------------------------------")
	number := 10
	number++
	fmt.Println(number)
	number--
	fmt.Println(number)
	// Finally Operation Unitary

}
