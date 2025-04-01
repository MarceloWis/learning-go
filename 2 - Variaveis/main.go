package main

import "fmt"

func main() {
	// 1 - Declaração de variáveis
	varClarations()
}

func varClarations() {
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constant string = "Constante"
	fmt.Println(constant)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
