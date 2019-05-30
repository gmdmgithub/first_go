package main

import (
	"fmt"
	"log"
	"time"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		time.Sleep(100 * time.Millisecond)
		return sum
	}
}

// StartTimer - make a time closure
func StartTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Now().Sub(t)
		log.Println(name, "took", d)
	}
}

func main() {
	took := StartTimer("main")
	defer took()
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
