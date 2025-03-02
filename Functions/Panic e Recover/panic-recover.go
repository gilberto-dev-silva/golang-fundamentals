package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Recover executado", r)
	}
}

func studentIsApproved(n1, n2 float64) bool {
	defer recoverExecution()
	media := (n1 + n2) / 2

	if media > 7 {
		return true
	} else if media < 7 {
		return false
	}
	panic("Media é exatamente 7")
}

func main() {
	approved := studentIsApproved(7, 7)
	fmt.Println(approved)

	fmt.Println("Pós execução")
}
