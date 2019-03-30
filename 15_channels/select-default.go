package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
	fmt.Println("Init function is called")
	time.Sleep(10 * time.Millisecond)
}

func selectDefault() {
	fmt.Println("\n#######################\nHi there select starts here!")
	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "Value 1"
	chan1 <- "Value 2"
	chan2 <- "Value 1"
	chan2 <- "Value 2"

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	default:
		fmt.Println("No goroutines available to send data", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
	fmt.Println("\n#######################\nHi there select finished here!")
}
