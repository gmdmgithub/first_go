package main

import (
	"fmt"
	"time"
)

// greet - just present greeting - but from channel
func greet(c chan string) {
	//channels have always let arrow notation <-  on the left of cannels read form channel,
	// on the right from channel write to channel
	fmt.Printf("Hi from channel: %s\n", <-c)
}

func channelSample() {

	//to be effective channels have to be created be make
	stringChannel := make(chan string)

	// if you do not create a goroutine an error will occur - comment the line below to check
	go greet(stringChannel)
	defer close(stringChannel)
	// time.Sleep(10 * time.Millisecond)
	stringChannel <- "Alex"

	fmt.Println("By from simple channel!\n############")
}

func greetMany(number int, c chan string) {

	for i := 0; i < number; i++ {
		fmt.Printf("Hi from blocking channel: %s\n", <-c)
	}
}

func chcekBlocking() {

	c := make(chan string)

	go greetMany(3, c)
	fmt.Println("Lte's send Alex")
	c <- "Alex"
	fmt.Println("Lte's send Marry")
	c <- "Marry"
	// go close(c)

	fmt.Println("Lte's send Kris - wait 1 sek ;-)")
	time.Sleep(1 * time.Second)
	c <- "Kris"

	time.Sleep(10 * time.Millisecond)

	fmt.Println("By from channel blocking!!")
}
