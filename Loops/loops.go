package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Looops in Go")
	i := 0

	for i < 5 {
		time.Sleep(time.Second)
		i++
		fmt.Println(i)
	}

	fmt.Println("----------------------------------------------------------------")
	for j := 0; j < 10; j += 5 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	fmt.Println("----------------------------------------------------------------")
	names := [3]string{"João", "Maria", "Pedro"}
	for index, name := range names {
		fmt.Printf("Name Position %d: %s\n", index, name)
	}

	fmt.Println("----------------------------------------------------------------")

	fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}
	for i, fruit := range fruits {
		fmt.Printf("Fruit Position %d: %s\n", i, fruit)
	}

	fmt.Println("----------------------------------------------------------------")

	for _, fruit := range fruits {
		go func(fruit string) {
			fmt.Println("Fruit:", fruit)
		}(fruit)
	}

	fmt.Println("----------------------------------------------------------------")

	for indice, word := range "Welcome" {
		fmt.Println("Character: ", indice, word)
	}

	fmt.Println("----------------------------------------------------------------")
	for _, letter := range "Welcome" {
		fmt.Println("Character: ", letter)
	}
	fmt.Println("----------------------------------------------------------------")

	for indice, letter := range "Welcome" {
		fmt.Printf("Character: %v - Index: %v\n", letter, indice)
	}

	fmt.Println("----------------------------------------------------------------")

	for indice, letter := range "Welcome" {
		fmt.Printf("Character: %v - Index: %v\n", indice, string(letter))
	}
}
