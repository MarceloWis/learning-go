package main

import "fmt"

func func1() {
	println("Func1")
}

func func2() {
	println("Func2")
}

func func3(n1, n2 float32) bool {
	fmt.Println("Func3")
	media := n1 + n2/2
	if media >= 7 {
		return true
	}
	return false
}

func main() {
	defer func1()
	func2()
	func2()
	func2()
	fmt.Println(func3(7, 8))
	func2()
	func2()
}
