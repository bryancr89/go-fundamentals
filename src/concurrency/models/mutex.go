package main

import (
	"runtime"
	"fmt"
	"sync"
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	//badPrints()
	//usingMutext()
	//usingChannels()
	usingChannels2()
}

func badPrints() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			go func(i, j int) {
				fmt.Printf("%d + %d = %d\n", i, j, i + j)
			}(i ,j)
		}
	}

	fmt.Scanln()
}

// Works slow
func usingMutext() {
	mutex := new(sync.Mutex)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex.Lock()
			go func(i, j int) {
				fmt.Printf("%d + %d = %d\n", i, j, i + j)
				mutex.Unlock()
			}(i ,j)
		}
	}

	fmt.Scanln()
}

// SHOULD NOT USE THIS
func usingChannels() {
	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <-  true
			go func(i, j int) {
				fmt.Printf("%d + %d = %d\n", i, j, i + j)
				<-mutex
			}(i ,j)
		}
	}

	fmt.Scanln()
}

func usingChannels2() {
	f, _ := os.Create("./log.txt")

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
			} else {
				f.Close()
				break
			}
		}
	}()
	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <-  true
			go func(i, j int) {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i + j)
				logCh <- msg
				<- mutex
			}(i ,j)
		}
	}

	fmt.Scanln()
}