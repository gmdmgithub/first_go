package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	log.Println("#### Hi there design patterns are essential in programing!")
	defer log.Println("---- By from design patterns are essential in programing!")
	creationalPattern()

}

func creationalPattern() {
	log.Println("## First is Creational pattern: Factory is an example")
	defer log.Println("-- By from Creational")

	kind := "my-writer"
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	// someRan := r1.Intn(500)
	rand.Seed(time.Now().UnixNano())
	someRan := rand.Intn(500)
	log.Printf("Random is %v", someRan)
	if someRan > 250 {
		kind = "stderr"
	}
	// Create writer and write some output
	writer, err := NewWriter(kind)
	if err != nil {
		log.Printf("Something went wrong %v", err)
		return
	}
	len, err := writer.Write([]byte("Hello world from Factory pattern\n"))
	if err != nil {
		log.Printf("Something went wrong %v", err)
		return
	}
	log.Printf("The length was %v", len)

	time.Sleep(500 * time.Millisecond) //it takes some time ;-)
}
