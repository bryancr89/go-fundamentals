package main

import (
	"fmt"
	"strings"
)

func main() {
	//basicChannel()
	//bufferedChannel()
	closeChannel()
}

func basicChannel() {
	ch := make(chan string, 1)
	ch <- "Hello"
	fmt.Println(<-ch)
}

func bufferedChannel() {
	phrase := "These are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}

func closeChannel() {
	phrase := "Closing channel message.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	//for {
	//	if msg, ok := <-ch; ok {
	//		fmt.Print(msg + " ")
	//	} else {
	//		break
	//	}
	//}

	// Simpler code
	for msg := range ch {
		fmt.Print(msg + " ")
	}
}