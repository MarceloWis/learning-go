package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	fmt.Println("------")
	a := 0
	for a < 10 {
		a++
	}
	fmt.Println(a)
	fmt.Println("------")

	nome := []string{"João", "Davi", "Lucas"}
	for i, valor := range nome {
		fmt.Println(i, valor)
	}
	fmt.Println("------")

	for i, valor := range "PALAVRA" {
		fmt.Println(i, string(valor))
	}
	fmt.Println("------")

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Marcelo",
			"ultimo":   "Wis",
		},
		"curso": {
			"nome": "Engenharia",
			"ano":  "2023",
		},
	}

	for chave, valor := range usuario2 {
		fmt.Println(chave, valor)
	}
	fmt.Println("------")
}
