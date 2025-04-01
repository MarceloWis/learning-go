package main

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	return somar(n1, n2), n1 - n2
}

func main() {
	soma := somar(10, 20)
	println(soma)

	var f = func() {
		println("Função f")
	}

	f()

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 20)
	println(resultadoSoma, resultadoSubtracao)

	resultadoSoma2, _ := calculosMatematicos(10, 20)
	println(resultadoSoma2)

}
