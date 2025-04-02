package main

func soma(numeros ...int) int {
	somaTotal := 0
	for _, numero := range numeros {
		somaTotal += numero
	}
	return somaTotal
}

func main() {
	println(soma(1, 2, 3, 4, 5))
}
