package main

import "fmt"

func dayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "1 - Sunday"
	case 2:
		return "2 - Monday"
	case 3:
		return "3 - Tuesday"
	case 4:
		return "4 - Wednesday"
	case 5:
		return "5 - Thursday"
	case 6:
		return "6 - Friday"
	case 7:
		return "7 - Saturday"
	default:
		return "Number invalid"
	}
}

func pets(number int) string {
	switch {
	case number == 1:
		return "Dog"
	case number == 2:
		return "Cat"
	case number == 3:
		return "Mouse"
	default:
		return "Number invalid"
	}
}

func cars(number int) string {
	var name string
	switch number {
	case 1:
		name = "Ferrari"
	case 2:
		name = "Mercedes"
	default:
		name = "Car invalid"
	}

	return name
}

func main() {
	fmt.Println("Switch")
	fmt.Println(dayOfTheWeek(1))
	fmt.Println(dayOfTheWeek(2))
	fmt.Println(dayOfTheWeek(3))
	fmt.Println(dayOfTheWeek(4))
	fmt.Println(dayOfTheWeek(5))
	fmt.Println(dayOfTheWeek(6))
	fmt.Println(dayOfTheWeek(7))
	fmt.Println(dayOfTheWeek(8))

	fmt.Println("\nSwitch with a variable")
	fmt.Println(pets(1))
	fmt.Println(pets(2))
	fmt.Println(pets(3))
	fmt.Println(pets(4))

	fmt.Println("\nSwitch with a variable cars")
	fmt.Println(cars(1))
	fmt.Println(cars(2))
	fmt.Println(cars(3))
}
