package main

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	println(fibonacci(uint(10)))
}
