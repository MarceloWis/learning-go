package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1

	v1++

	fmt.Println(v1, v2)

	fmt.Println("-----")

	var v3 int = 10
	var v4 *int = &v3

	v3++

	fmt.Println(v3, *v4) // desreferenciando o ponteiro
	fmt.Println(v3, v4)  // endereço de memória
}
