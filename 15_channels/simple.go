package main

import (
	"fmt"
)

// greet - just present greeting - but from channel
func greet(c chan string) {
	//channels have always let arrow notation <-  on the left of cannels read form channel,
	// on the right from channel write to channel
	fmt.Printf("Hi from channel: %s\n", <-c)
}

func channelSample() {

	//to be effective channels have to be created be make
	strinChannel := make(chan string)

	// if you do not create a goroutine an error will occur - comment the line below to check
	go greet(strinChannel)
	defer close(strinChannel)
	// time.Sleep(10 * time.Millisecond)
	strinChannel <- "Alex"

	fmt.Println("By from channel")
}
