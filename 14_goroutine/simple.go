package main

import (
	"fmt"
	"time"
)

func sayHello(isSleeping bool) {
	fmt.Printf("Hi there! %t\n", isSleeping)
}

func simpleGoroutine(isSleeping bool) {
	fmt.Println("main execution started")

	// create goroutine
	go sayHello(isSleeping)

	//test if sayHello will affect execution
	if isSleeping {
		// schedule another goroutine
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("main execution stopped and sleeping was: %t\n", isSleeping)
}
