package main

import "fmt"

func Array() {
	// Arrays
	var a1 [5]int
	a1[0] = 1
	a1[1] = 2
	a1[2] = 3
	a1[3] = 4
	a1[4] = 5
	fmt.Println(a1)

	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a2)

	a3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a3)
}

func Slices() {
	// Slices
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)

	s1 = append(s1, 6)
	fmt.Println(s1)
	a1 := [5]int{1, 2, 3, 4, 5}
	s2 := a1[1:3]
	fmt.Println(s2)

}

func main() {
	Array()
	fmt.Println("-----")
	Slices()
}
