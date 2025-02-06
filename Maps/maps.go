package main

import "fmt"

func main() {
	fmt.Println("Maps")

	map1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}
	fmt.Println(map1)
	fmt.Println(map1["one"])
	fmt.Println(map1["two"])
	fmt.Println(map1["three"])
	fmt.Println(map1["four"])
	fmt.Println(map1["five"])

	map2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(map2)
	delete(map2, "nome")
	fmt.Println(map2)
}
