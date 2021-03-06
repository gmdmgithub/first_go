package main

import (
	"fmt"
	"log"
	"time"
)

var counter = 1 // manipulate parallel number of executions

var semaphoreChan = make(chan struct{}, counter)

func doTheJob(index int) {

	semaphoreChan <- struct{}{} //block while its full

	go func() {
		defer func() {
			<-semaphoreChan //read to release
		}()
		log.Println(index + 1)
		time.Sleep(1 * time.Second) // simulate long lasting task
	}()

}

func main() {
	fmt.Println("Hi semaphore here ----------")
	defer fmt.Println("Bye semaphore  ----------")

	for i := 0; i < 20; i++ {
		doTheJob(i)
	}

	benchmarkMux(40, 20)
}
