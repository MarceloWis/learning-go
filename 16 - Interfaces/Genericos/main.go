package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(1)

	result := map[string]interface{}{
		"nome":  "Alejandro",
		"idade": 22,
	}

	generica(result)

}
