package main

import (
	"log"
	"math/rand"
	"time"
)

// MyDecorFunc - blueprint
type MyDecorFunc func(count int) (ok bool)

func structuralPattern() {
	log.Println("## Second is Structural pattern: Decorator pattern is an example")
	defer log.Println("-- By from Structural")
	//in real go programing the http handler func is a good example like isAuthorized, loggin ..
	DecorFunc(MyFunc)
}

// MyFunc - the same footprint as blueprint func, just for test
func MyFunc(c int) (success bool) {

	log.Printf("How is? %d", c)
	// rundom success
	return c > 10
}

// DecorFunc - func to cover (embrace) the MyFunc
func DecorFunc(someFunc MyDecorFunc) {

	log.Printf("I just started %v", time.Now())
	defer log.Printf("And finished  %v", time.Now())
	time.Sleep(1 * time.Second) //do hard work
	log.Println("Hard work has been done!")

	// heart!! someFunc passed to the func is called to get decor
	if ok := someFunc(rand.Intn(20)); ok {
		log.Println("Number was bigger than 10 ;-)")
		return
	}

	log.Println("Not successful, but why who knows")
}
