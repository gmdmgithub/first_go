package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	fmt.Println("########### Hi there - we will try waiting groups to synchronise goroutine")
	defer fmt.Println("########### By from waiting groups to synchronise goroutine")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() { //first long lasted function in goroutine
		defer wg.Done()

		time.Sleep(1 * time.Second) //some long lasting task
		log.Println("First long lasted finished")

	}()
	// the main thread is doing something meanwhile
	time.Sleep(300 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second) //some other long lasted
		log.Println("Second long lasted finished")
	}()

	wg.Wait()
}
