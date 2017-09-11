package main

import "fmt"

func main() {
	basicChannel()
}

func basicChannel() {
	ch := make(chan string, 1)
	ch <- "Hello"
	fmt.Println(<-ch)
}