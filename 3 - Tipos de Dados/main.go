package main

import (
	"errors"
	"fmt"
)

func main() {
	Inteiros()
	Reais()
	Strings()
	Bool()
	Error()
}

func Inteiros() {
	// quanto cada int suporta
	// int8 -> -128 a 127
	// int16 -> -32768 a 32767
	// int32 -> -2147483648 a 2147483647
	// int64 -> -9223372036854775808 a 9223372036854775807

	// int8, int16, int32, int64
	var numero1 int16 = 100
	// var numero2 int8 = 1000000000000 // cannot use 1000000000000 (untyped int constant) as int8 value in variable declaration
	fmt.Println(numero1)

	// int -> define o tamanho da variável de acordo com o SO (32 ou 64 bits)

	// uint8, uint16, uint32, uint64
	// uint -> 0 a 4294967295
	// uint8 -> 0 a 255
	// uint16 -> 0 a 65535
	// uint32 -> 0 a 4294967295
	// uint64 -> 0 a 18446744073709551615
	var numero2 uint16 = 1000
	fmt.Println(numero2)

	// alias

	var numero3 rune = 12345 // int32
	fmt.Println(numero3)

	var numero4 byte = 123 // uint8
	fmt.Println(numero4)
}

func Reais() {
	// float32, float64
	// float32 -> 1.401298464324817070923729583289916131280e-45 a 3.40282346638528859811704183484516925440e+38
	// float64 -> 4.940656458412465441765687928682213723651e-324 a 1.797693134862315708145274237317043567981e+308
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	// float -> float32 ou float64 de acordo com o SO
	var numeroReal3 float64 = 123.45
	fmt.Println(numeroReal3)
}

func Strings() {
	// string
	var texto string = "Texto"
	fmt.Println(texto)

	// char -> não existe
	var char byte = 'T'
	fmt.Println(char) // 84 (valor na tabela ASCII)

	var texto2 string   // Variável string iniciada vazia
	fmt.Println(texto2) // string vazia

}

func Bool() {
	var bool1 bool = true
	fmt.Println(bool1)

	var bool2 bool
	fmt.Println(bool2)

}

func Error() {
	var error1 error // Variável error iniciada vazia nil
	fmt.Println(error1)

	var error2 error = errors.New("Erro")
	fmt.Println(error2)
}
