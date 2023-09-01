package main

import (
	"fmt"
	"time"
)

func Hello() {
	for {
		time.Sleep(time.Second)
		fmt.Println("Hello, world!")
	}
}
func main() {
	chan_A := make(chan bool)
	chan_B := make(chan bool)
	go func() {
		for {
			<-chan_A
			fmt.Println("A")
			chan_B <- true
		}
	}()
	go func() {
		for {
			<-chan_B
			fmt.Println("B")
			chan_A <- true
		}
	}()
	chan_A <- true
	time.Sleep(999 * time.Second)
}
