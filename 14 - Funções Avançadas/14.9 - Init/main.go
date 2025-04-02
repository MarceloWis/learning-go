package main

import "fmt"

var n int

func main() {
	fmt.Println("Main", n) // 10
}

func init() {
	fmt.Println("Init")
	n = 10
}
