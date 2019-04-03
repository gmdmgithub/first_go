package main

import (
	"fmt"
	"sync"
)

var i int

// makeWork adds value to i variable - without mutex often may modify the same value on heap - check by changing useMtex
// mutex (lock) cause its accessible obly for one goroutine
func makeWork(w *sync.WaitGroup, m *sync.Mutex, useMutex bool) {

	if useMutex {
		m.Lock()
	}
	i = i + 1

	if useMutex {
		m.Unlock()
	}

	w.Done()
}

func perfromMutex() {
	fmt.Println("\n#######################\nHi there,perfromMutex - starts here!")

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for j := 0; j < 1500; j++ {
		wg.Add(1)
		go makeWork(&wg, &mutex, true)
	}

	wg.Wait()

	fmt.Printf("Finally the value of i is: %d\n", i)

	fmt.Println("\n#######################\nHi there,perfromMutex - ends here!")

}
