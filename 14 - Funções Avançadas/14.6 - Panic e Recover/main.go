package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução")

	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func calculoDeMedia(n1, n2 float32) bool {
	defer recuperarExecucao()
	fmt.Println("calculoDeMedia")
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é 6")
}

func main() {
	fmt.Println(calculoDeMedia(6, 6))
}
