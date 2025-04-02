package main

import (
	"fmt"
	"time"
)

func escrever(text string, channel chan string) {
	for i := 0; i < 5; i++ {

		channel <- text
		time.Sleep(time.Second)
	}
	close(channel)
}

func main() {
	channel := make(chan string)
	go escrever("Hello World", channel)
	go escrever("Hello World2", channel)

	// for {
	// 	result, aberto := <-channel
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(result, aberto)
	// }
	for result := range channel {
		fmt.Println(result)
	}

}
