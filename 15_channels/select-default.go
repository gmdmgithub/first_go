// select is used where you needed only one conditions to be true
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

func serviceGoroutine(c chan string) {
	c <- "My name is Alex"
}

func selectWitGoroutine() {

	fmt.Println("\n#######################\nHi there, selectWitGoroutine starts here!")

	// var channel chan string
	channel := make(chan string)

	go serviceGoroutine(channel)

	time.Sleep(10 * time.Millisecond)

	select {
	case res := <-channel:
		fmt.Printf("Channel has a value: %v\n", res)

	default:
		fmt.Println("Channel has no value")
	}

	fmt.Println("\nHi there, selectWitGoroutine finish here!\n#######################")
}

func serviceWithDelay(c chan string, delay int) {

	time.Sleep(time.Duration(delay) * time.Second)
	c <- fmt.Sprintf("Alex here! - simulate %d sec delay", delay)
}

func playingWithTimeout() {

	fmt.Println("\n#######################\nHi there, playingWithTimeout starts here! - very useful!!")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go serviceWithDelay(channel1, 3)
	go serviceWithDelay(channel2, 1)

	select {
	case res := <-channel1:
		fmt.Printf("Good job channel1 you are on time! You are: %s\n", res)
	case res := <-channel2:
		fmt.Printf("Good job channel2 you are on time! You are: %s\n", res)
	case <-time.After(2 * time.Second):
		fmt.Println("No answer time-out and bye bye!")
	}

	fmt.Println("\n#######################\nHi there, playingWithTimeout ends here! - bye!")
}
