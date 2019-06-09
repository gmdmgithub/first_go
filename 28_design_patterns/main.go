package main

import (
	"log"
)

func main() {

	log.Println("#### Hi there design patterns are essential in programing!")
	defer log.Println("---- By from design patterns are essential in programing!")
	creationalPattern()
	structuralPattern()
	behavioralPattern()
}
