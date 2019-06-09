package main

import (
	"errors"
	"log"
)

type Iterator struct {
	tasks    []string
	position int
}

func behavioralPattern() {
	log.Println("## Third is Behavioral pattern: Iterator pattern is an example (io.Read - or result in db query)")
	log.Println("## Accesses elements sequentially, without exposing underlying representation")
	defer log.Println("-- By from Behavioral")

	tasks := GetExamples()
	// infinite loop can be used
	for {
		i, t, err := tasks.Next()
		if err == ErrEOF {
			log.Printf("Done")
			break
		}
		if err != nil {
			log.Fatalf("Unknown error: %s", err)
		}
		log.Printf("Task %d: %s\n", i, t)
	}

}

func GetExamples() Iterator {
	return Iterator{
		tasks: []string{
			"say hello",
			"say goodbye",
		},
	}
}

// ErrEOF signals a graceful end of iteration
var ErrEOF = errors.New("EOF")

// Next will return the next task in the slice, or ErrEOF
func (t *Iterator) Next() (int, string, error) {
	t.position++
	if t.position > len(t.tasks) {
		return t.position, "", ErrEOF
	}
	return t.position, t.tasks[t.position-1], nil
}
