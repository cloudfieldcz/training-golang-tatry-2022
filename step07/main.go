package main

import (
	"fmt"
	"time"
)

var cancelLoop chan bool

func asyncLoop() {
	for {
		select {
		case <-cancelLoop:
			fmt.Println("task canceled!")
			return
		case <-time.After(600 * time.Millisecond):
			fmt.Println("task timeout")
		}
	}
}

func main() {
	cancelLoop = make(chan bool)
	fmt.Println("start ..")

	// goroutine
	go asyncLoop()

	time.Sleep(3 * time.Second)
	fmt.Println("send cancel ..")
	// and there is the cancel of async func!
	cancelLoop <- true
	time.Sleep(1 * time.Second)
	fmt.Println("end ..")
}
