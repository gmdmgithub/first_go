// Wait group is used as opposit to select where you needed only one conditions to be true,
// but with wait you need all conditions to be true in order to unblock the main goroutine

package main

import (
	"fmt"
	"sync"
	"time"
)

type someStruct struct {
	name string
	age  int
}

func serveMe(wg *sync.WaitGroup, sS *[]someStruct, number int) {
	defer wg.Done() //here we decrement a counter
	fmt.Printf("Hi there waitgroup is called %d time\n", number)
	s := someStruct{
		"Alex",
		number,
	}
	time.Sleep(1 * time.Second)
	*sS = append(*sS, s)
	time.Sleep(1 * time.Second)
}

func waitForAllGoroutines() {

	fmt.Println("\n#######################\nHi there, wiat Group - as opposite to select starts here!")
	defer fmt.Println("\n#######################\nHi there, wait Group - as opposite to select ends here!")

	var wg sync.WaitGroup

	var res []someStruct

	for i := 1; i <= 5; i++ {
		wg.Add(1) //here we increment counter - number how many!!
		go serveMe(&wg, &res, i)
	}

	wg.Wait() //here we block

	fmt.Printf("End job and show results: %+v", res)
}
