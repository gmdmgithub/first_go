package main

import (
	"fmt"
	"runtime"
)

func doSomething(c chan int) {

	for i := 0; i < 3; i++ {
		chanVal := <-c
		calc := chanVal * 3
		fmt.Printf("Multiply channel %d by 3 is :%d \n", chanVal, calc)
	}
}

func blockingChannel() {

	fmt.Println("blockingChannel: Checking channel blocking and capacity")

	c := make(chan int, 3)
	go doSomething(c)
	defer close(c)
	c <- 3
	c <- 5
	c <- 7

	c <- 11 //this should block - comment to see

	fmt.Printf("blockingChannel: capacity %d, length %d and value: %d", cap(c), len(c), <-c)

	go doSomething(c)

	c <- 12
	c <- 14
	c <- 16

	c <- 20 //block again

	fmt.Println("number of goroutine?", runtime.NumGoroutine())

	fmt.Printf("blockingChannel: Finished ...")
}
