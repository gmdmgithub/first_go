package main

import "fmt"

func convertToRead(c <-chan string) {
	//now channel is read-only
	fmt.Println("Read chaannel welcome! Just converted to read only", <-c)
}

func writeChannel(c chan<- int) {
	fmt.Printf("Tel me abour yourself %T\n", c)
}

func uniChannelPalyground() {
	fmt.Println("\n#######################\nHi there unichannel here!")

	ordinaryChannel := make(chan string)

	reciveChannel := make(<-chan int)
	senderChannel := make(chan<- int)

	// impossible - syntax error
	// reciveChannel <- 2
	fmt.Printf("What type is reciver? %T\n", reciveChannel)

	go writeChannel(senderChannel)

	go convertToRead(ordinaryChannel)

	defer close(ordinaryChannel)
	// defer close(reciveChannel)// receive only channel cannot be closed
	defer close(senderChannel)

	ordinaryChannel <- "My name is Alex"

	// senderChannel <- 2 //correct - channel may get data

	fmt.Println("By from unichannel")
}
