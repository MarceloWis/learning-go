package main

func main() {
	numero := 10

	if numero > 15 {
		println("Maior que 15")
	} else {
		println("Menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		println("Número é maior que zero")
	}
}
