package main

import "fmt"

func convertToRead(c <-chan string) {
	//now channel is read-only
	fmt.Println("Read channel welcome! Just converted to read only", <-c)
}

func writeChannel(c chan<- int) {
	fmt.Printf("Tel me about yourself %T\n", c)
}

func uniChannelPlayground() {
	fmt.Println("\n#######################\nHi there unichannel here!")

	ordinaryChannel := make(chan string)

	reciveChannel := make(<-chan int)
	senderChannel := make(chan<- int)

	// impossible - syntax error
	// receiveChannel <- 2
	fmt.Printf("What type is reciver? %T\n", reciveChannel)

	go writeChannel(senderChannel)

	go convertToRead(ordinaryChannel)

	defer close(ordinaryChannel)
	// defer close(receiveChannel)// receive only channel cannot be closed
	defer close(senderChannel)

	ordinaryChannel <- "My name is Alex"

	// senderChannel <- 2 //correct - channel may get data

	fmt.Println("By from unichannel")
}
