package main

import (
	"fmt"
	"time"
)

func main() {

	go loop()
	// block indefinitely. We're reading from an anonymous channel that never has anything sent to it
	<-make(chan int)
}

func loop() {
	for {
		fmt.Println("Printing")
		time.Sleep(1 * time.Second)
	}
}
