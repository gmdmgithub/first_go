package main

import (
	"fmt"
	"sync"
	"time"
)

// worker than make squares
func sqrWorker(tasks <-chan int, results chan<- int, id int, wg *sync.WaitGroup) {
	for num := range tasks {
		time.Sleep(10 * time.Millisecond) // simulating blocking/performing task
		fmt.Printf("[worker no: %v] Sending result by worker %v\n", id, id)
		results <- num * num
	}
	wg.Done()
}

func workerPoolPlaygound() {
	fmt.Println("\n#######################\nHi there, worker pool - starts here!")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

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

	//stopinng here all rorker will finish itrs job
	// time.Sleep(5 * time.Second)

	//similar effect (above sleep) but effective
	wg.Wait()

	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")
}
