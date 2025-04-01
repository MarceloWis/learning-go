package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, world! Main")

	auxiliar.Escrever()
	auxiliar.Escrever2()
	err := checkmail.ValidateFormat("marcelowiscgmail.com")
	fmt.Println(err)
}
