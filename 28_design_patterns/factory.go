package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func creationalPattern() {
	log.Println("## First is Creational pattern: Factory is an example")
	defer log.Println("-- By from Creational")

	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	// someRan := r1.Intn(500)
	rand.Seed(time.Now().UnixNano())
	someRan := rand.Intn(500)
	log.Printf("Random is %v", someRan)
	kind := "my-writer"
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

// NewWriter - Creational: Factory pattern example
func NewWriter(kind string) (io.Writer, error) {
	switch kind {
	case "my-writer":
		return MyWriter{}, nil
	case "stderr":
		return os.Stderr, nil
	default:
		return nil, fmt.Errorf("Invalid writer was sent: %s", kind)
	}
}

// MyWriter - example struct to show factory pattern
type MyWriter struct{}

func (w MyWriter) Write(p []byte) (n int, err error) {
	log.Printf("My Writer: %s", p)
	return len(p), nil
}
