package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

// makeWork adds value to i variable - without mutex often may modify the same value on heap - check by changing useMtex
// mutex (lock) cause its accessible only for one goroutine
func makeWork(w *sync.WaitGroup, m *sync.Mutex, useMutex bool) {

	if useMutex {
		m.Lock()
	}
	time.Sleep(500 * time.Millisecond)
	i = i + 1

	if useMutex {
		m.Unlock()
	}
	time.Sleep(500 * time.Millisecond)

	w.Done()
}

func perfromMutex() {
	fmt.Println("\n#######################\nHi there,perfromMutex - starts here!")
	defer fmt.Println("\n#######################\nHi there,perfromMutex - ends here!")

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for j := 0; j < 1500; j++ {
		wg.Add(1)
		go makeWork(&wg, &mutex, true)
	}

	wg.Wait()

	fmt.Printf("Finally the value of i is: %d\n", i)

}
