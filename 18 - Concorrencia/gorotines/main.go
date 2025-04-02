package main

func escrever(text string) {
	for {
		println(text)
	}
}

func main() {
	go escrever("Hello World")
	escrever("Programando em Go")
}
