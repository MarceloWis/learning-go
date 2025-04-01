package main

import "fmt"

func main() {
	// Aritméticos
	// + - * / %
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 1 / 2
	restoDaDivisao := 1 % 2
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	// Atribuição
	var variavel string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel, variavel2)

	// Relacionais
	// == != > < >= <=
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)

	// Lógicos
	// && || !
	fmt.Println(true && true)
	fmt.Println(true || true)
	fmt.Println(!true)

	// Unários
	// ++ --
	x := 1
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	// Ternário em Go
	var texto string
	if x > 10 {
		texto = "Maior que 10"
	} else {
		texto = "Menor que 10"
	}
	fmt.Println(texto)

}
