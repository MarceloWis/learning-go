package main

import "fmt"

type User struct {
	nome     string
	idade    uint8
	endereco Address
}

type Address struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Hello, world!")
	var u User
	u.nome = "Fulano"
	u.idade = 18
	fmt.Println(u)

	u2 := User{"Ciclano", 20, Address{"Rua A", 10}}
	fmt.Println(u2)

	u3 := User{idade: 30}
	fmt.Println(u3)

}
