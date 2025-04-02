package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Marcelo",
		"sobrenome": "Wis",
	}

	fmt.Println(usuario)

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

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
		"ano":  "2023",
	}
	fmt.Println(usuario2)

}
