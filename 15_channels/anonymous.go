package main

import "fmt"

func playAnonymousGorutine() {

	fmt.Println("\n#######################\nHi there anonymous here!")

	c := make(chan string)

	// launch anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello from anonymous" + <-c + "!")
	}(c)

	c <- "Alex"
	fmt.Println("\n#######################\nBy anonymous!")
}
