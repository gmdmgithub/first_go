package main

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

// Func - func type with matching signature
type Func func(attempt int) (retry bool, err error)

// MaxRetries - global value of retries
const MaxRetries = 5

func main() {
	err := Try(doCoupleTimes)
	if err != nil {
		log.Fatalln("error:", err)
	}
}

func doCoupleTimes(attempt int) (tryAgain bool, err error) {

	//maybe random true false?
	e := doSomeLongTermJob()

	return true, e
}

func doSomeLongTermJob() (err error) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return nil
}

// Try - function to retrail func execution
func Try(fn Func) error {
	var err error
	var tryAgain bool
	attempt := 1
	for {
		tryAgain, err = fn(attempt)
		log.Printf("try more? %t and err %v, attempt no %d", tryAgain, err, attempt)
		if !tryAgain && err == nil {
			break
		}
		attempt++
		if attempt > MaxRetries {
			return errors.New("Exceed number attempts")
		}
	}
	return err
}
