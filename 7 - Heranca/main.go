package main

import "fmt"

type (
	Pessoa struct {
		nome      string
		sobrenome string
		idade     int
		altura    uint8
	}
	Estudante struct {
		Pessoa
		curso     string
		faculdade string
		matricula int
	}
)

func main() {
	p1 := Pessoa{
		nome:      "Lucas",
		sobrenome: "Mendes",
		idade:     25,
		altura:    180,
	}

	fmt.Println(p1)

	e1 := Estudante{
		Pessoa:    p1,
		curso:     "Engenharia",
		faculdade: "USP",
		matricula: 123456,
	}
	fmt.Println(e1)
}
