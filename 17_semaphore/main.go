package main

import (
	"fmt"
	"log"
	"time"
)

var countetr = 1 // manipulate paralel namber of executions

var semaphoreChan = make(chan struct{}, countetr)

func doTheJob(index int) {

	semaphoreChan <- struct{}{} //block while its full

	go func() {
		defer func() {
			<-semaphoreChan //read to release
		}()
		log.Println(index + 1)
		time.Sleep(3 * time.Second) // simulate long lasting task
	}()

}

func main() {
	fmt.Println("Hi semaphore here ----------")
	defer fmt.Println("Bye semaphore  ----------")

	for i := 0; i < 20; i++ {
		doTheJob(i)
	}
}
