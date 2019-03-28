package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Goroutine playground is here!")

	// Anonymous goroutine
	go func() {

		fmt.Println("I have no name but I'm working")
	}()

	time.Sleep(20 * time.Millisecond)

	isSleeping := false
	simpleGoroutine(isSleeping)
	simpleGoroutine(!isSleeping)
}
