package main

import (
	"log"
	"time"
)

func playAnonymousGoroutine() {
	start := time.Now()
	log.Println("\n#######################\nHi there anonymous here!")
	defer log.Printf("\n#######################\nBy anonymous! takes %v\n", time.Since(start))
	c := make(chan string)

	// launch anonymous goroutine
	go func(c chan string) {
		time.Sleep(1 * time.Second) //do something expensive
		log.Println("Hello from anonymous" + <-c + "!")
	}(c)

	log.Printf("Time? %v\n", time.Since(start))
	c <- "Alex" //this will wait till go routine end

}
