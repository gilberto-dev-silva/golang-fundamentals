package main

import (
	"fmt"
	"modulo/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing first module in Go!")
	assistant.Writing()

	err := checkmail.ValidateFormat("rereewrwe")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("E-mail invalid")
	}
}
