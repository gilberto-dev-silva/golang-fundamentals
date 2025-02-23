package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello, anonymous function!")
	}()

	func(text string) {
		fmt.Println(text)
	}("Passando parametro para anonymous function")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
