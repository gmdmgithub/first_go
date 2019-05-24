package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// MutexBlocker - mutex struct
type MutexBlocker struct {
	locker *sync.Mutex
}

//NewMutexBlocker - constructor
func NewMutexBlocker() *MutexBlocker {
	return &MutexBlocker{locker: &sync.Mutex{}}
}

func (m *MutexBlocker) Read() interface{} {
	m.locker.Lock()
	defer m.locker.Unlock()
	time.Sleep(60 * time.Millisecond)
	log.Printf("Read done")
	return "good job"
}

func (m *MutexBlocker) Write(l interface{}) error {

	m.locker.Lock()
	defer m.locker.Unlock()
	time.Sleep(50 * time.Millisecond)
	log.Printf("Write done %v", l)
	return nil
}

func benchmarkMux(workers int, workload int) {
	log.Println("benchmarkMux start")
	defer log.Println("benchmarkMux end")
	m := NewMutexBlocker()
	wg := &sync.WaitGroup{}

	writer := func() {
		for i := 0; i < workload; i++ {
			m.Write(fmt.Sprintf("Test %d", i))
		}
		wg.Done()
	}

	reader := func() {
		for i := 0; i < workload; i++ {
			m.Read()
		}
		wg.Done()
	}

	for i := 0; i < workload; i++ {
		wg.Add(1)
		go writer()
		wg.Add(1)
		go reader()
	}

	wg.Wait()

}
