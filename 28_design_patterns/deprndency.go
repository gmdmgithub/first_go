package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Creational: Dependency Injection (Constructor) pattern
func dependencyInjection() {
	log.Println("## Creational: Dependency Injection (Constructor) pattern")
	defer log.Println("-- By from Dependency Injection")
	s := NewMyService(os.Stderr)
	s.WriteHello("world")
}

// MyService - basic struct
type MyService struct {
	writer io.Writer
}

// NewMyService - constructor of the struct
func NewMyService(writer io.Writer) MyService {
	return MyService{
		writer: writer,
	}
}

// WriteHello - sample method
func (s *MyService) WriteHello(m string) {
	log.SetOutput(s.writer)
	fmt.Fprintf(s.writer, "Hello %s\n", m)
}
