package main

import (
	"fmt"
	"time"
)

func putToChannel(c chan int, num int) {

	for i := 0; i < num; i++ {
		c <- i * 2
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(1 * time.Second)
	c <- 1000 //joke ;-)

	close(c)
}

// procesInfiniteWait - example how to use
func procesInfiniteWait() {
	fmt.Println("Welcome to the infinite loop with channels")

	c := make(chan int)

	go putToChannel(c, 20)

	//waiting for answers ...
	for {
		val, ok := <-c
		if ok == false {
			fmt.Println("Channel is closed :(")
			break
		} else {
			fmt.Printf("We get the channel value: %d\n", val)
		}
	}

	fmt.Println("By by from infinite loop with channels")
}
