package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo primeiro modulo em Go!")
	auxiliar.Escrevendo()

	err := checkmail.ValidateFormat("rereewrwe")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("E-mail valido")
	}
}