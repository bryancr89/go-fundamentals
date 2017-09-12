package main

import "fmt"

func main() {
	ch := make(chan int)
	go generate(ch)

	for {
		prime := <- ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		fmt.Println("omg")
		ch = ch1
	}
}

func hello() {
	fmt.Println("HelloWorld")
}

func generate(ch chan int) {
	for i := 2; i<15; i++ {
		fmt.Println("hola ",i)
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		fmt.Println("hols",i, prime)
		if i % prime != 0 {
			out <- i
		}
	}
}