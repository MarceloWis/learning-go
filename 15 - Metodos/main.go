package main

import "fmt"

type Usuario struct {
	nome  string
	idade uint8
}

func (u Usuario) salvar() {
	fmt.Println("Salvando os dados do usuário", u.nome, u.idade)
}

func (u *Usuario) fazerAniversario() {
	fmt.Println("Fazendo aniversário do usuário", u.nome)
	u.idade++
}

func main() {
	usuario := Usuario{"Luiz", 20}
	usuario.fazerAniversario()
	usuario.salvar()
}
