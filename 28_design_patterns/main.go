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
	structuralPattern()
}

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

type MyDecorFunc func(count int) (ok bool)

func structuralPattern() {
	log.Println("## Second is Structural pattern: Decorator pattern is an example")
	defer log.Println("-- By from Structural")
	//in real go programing the http handler func are good examples like isAuthorized, loggin ..
	DecorFunc(MyFunc)
}

func MyFunc(c int) (success bool) {

	log.Printf("How is? %d", c)
	return c > 10
}

func DecorFunc(someFunc MyDecorFunc) {

	log.Printf("I just started %v", time.Now())
	defer log.Printf("And finished  %v", time.Now())
	time.Sleep(1 * time.Second) //do hard work
	log.Println("Hard work has been done!")
	if ok := someFunc(rand.Intn(20)); ok {
		log.Println("Number was bigger than 10 ;-)")
	}
}
