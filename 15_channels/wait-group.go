// Wait group is used as opposit to select where you needed only one conditions to be true,
// but with wait you need all conditions to be true in order to unblock the main goroutine

package main

import (
	"fmt"
	"sync"
	"time"
)

func serveMe(wg *sync.WaitGroup, number int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Hi there waitgroup is called %d time\n", number)
	wg.Done() //here we decrement a counter

}

func waitForAllGoroutines() {

	fmt.Println("\n#######################\nHi there, wiat Group - as opposite to select starts here!")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) //here we increment counter - numer how many!!
		go serveMe(&wg, i)
	}

	wg.Wait() //here we block

	fmt.Println("\n#######################\nHi there, wiat Group - as opposite to select ednds here!")
}
