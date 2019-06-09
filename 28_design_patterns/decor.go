package main

import (
	"log"
	"math/rand"
	"time"
)

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
