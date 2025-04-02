package main

import "sync"

func escrever(text string) {
	for i := 0; i < 5; i++ {
		println(text)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Hello World")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
