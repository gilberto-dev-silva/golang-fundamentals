package assistant

import "fmt"

// Escrevendo prints a message indicating it is being called from the auxiliar package.
func Writing() {
	fmt.Println("Writing from the helper package")
	toWrite()
}
