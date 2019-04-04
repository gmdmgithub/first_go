package main

import (
	"fmt"
	"time"
)

// TimeTracker - universal function for time tracking
func TimeTracker(start time.Time) {
	fmt.Printf("It took %v\n", time.Since(start))
}

func main() {
	fmt.Println("--------Hi we have a job!")
	defer fmt.Println("Job is done! -------------------")

	timeTrackerWay()
	traditionTimeTracking()
}

func timeTrackerWay() {

	defer TimeTracker(time.Now()) //remember defer go on heap and have lifo que

	time.Sleep(3 * time.Second) //do a hard job
}

func traditionTimeTracking() {

	start := time.Now()

	//do something
	time.Sleep(1 * time.Second)

	fmt.Println(time.Since(start))

}
