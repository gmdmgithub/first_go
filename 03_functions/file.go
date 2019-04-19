package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// FILE I/O
func fileOperation() {
	fmt.Println("###### File operations ######")

	// Create a file
	file, err := os.Create("samp.txt")

	// Close the file
	defer file.Close()

	// Output any errors
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("OK file created")
	}

	// Write a string to the file
	file.WriteString("Lets write some text to the file")

	// Try to open the file
	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Convert into a string
	readString := string(stream)

	fmt.Println(readString)
}
