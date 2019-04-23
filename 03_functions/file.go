package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// FILE I/O
func fileOperation() {
	fmt.Println("###### File operations start ######")
	defer fmt.Println("###### File operations End ######")

	// Create a file
	file, err := os.Create("samp.txt")

	// Close the file
	defer file.Close()

	// Output any errors
	if err != nil {
		log.Fatal(err) //fatal program exit!
	}

	log.Print("OK file created")

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

func fileStatistics() {
	fs, err := os.Stat("samp.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File does not exits - maybe some action?")
		}
		log.Fatal(err) //fatal program exit!
	}

	fmt.Printf("File data are: %+v\n", fs)
}

func checkSum() {

	// Get bytes from file
	data, err := ioutil.ReadFile("samp.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Hash the file and output results
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
}
