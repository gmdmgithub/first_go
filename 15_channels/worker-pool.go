package main

import (
	"fmt"
	"sync"
	"time"
)

type myData struct {
	name string
	age  int
}

type result struct {
	surname string
	old     bool
}

// worker than make squares
func sqrWorker(tasks <-chan int, results chan<- result, id int, wg *sync.WaitGroup) {
	for num := range tasks {
		time.Sleep(500 * time.Millisecond) // simulating blocking/performing task
		fmt.Printf("[worker no: %v] Sending result by worker %v\n", id, id)
		// old := tasks.age > 21
		fmt.Printf("num is %d\n", num)
		res := result{fmt.Sprintf("Alex %d", num), num > 2}
		results <- res
	}
	wg.Done()
}

func workerPoolPlayground() {
	fmt.Println("\n#######################\nHi there, worker pool - starts here!")

	tasks := make(chan int, 10)
	results := make(chan result, 10)

	var wg sync.WaitGroup

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker(tasks, results, i, &wg)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	//stopping here all broker will finish it's job
	// time.Sleep(5 * time.Second)

	//similar effect (above sleep) but effective
	wg.Wait()

	// receiving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}
	fmt.Println("[main] main() stopped")
}
