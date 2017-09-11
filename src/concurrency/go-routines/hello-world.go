package main

import "time"

func main() {
	helloWorldWithSleep()
}

func helloWorldWithSleep() {
	go func() {
		println("Hello")
	}()

	go func() {
		println("Go")
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

