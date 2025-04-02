package main

func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		println(texto)
	}

	return funcao
}

func main() {
	funcao := closure()
	funcao()
}
