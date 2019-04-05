package main

import (
	"fmt"
	"time"
)

func waitThenSend(i int, value string) chan int {
	channel := make(chan int)

	go func() {
		time.Sleep(time.Duration(i) * time.Second) //simulate long time job
		channel <- i
	}()

	return channel
}

func main() {
	fmt.Println(<-waitThenSend(4, "Alex"))
}
