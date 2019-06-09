package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// NewWriter - Creational: Factory pattern example
func NewWriter(kind string) (io.Writer, error) {
	switch kind {
	case "my-writer":
		return MyWriter{}, nil
	case "stderr":
		return os.Stderr, nil
	default:
		return nil, fmt.Errorf("Invalid writer was sent: %s", kind)
	}
}

// MyWriter - example struct to show factory pattern
type MyWriter struct{}

func (w MyWriter) Write(p []byte) (n int, err error) {
	log.Printf("My Writer: %s", p)
	return len(p), nil
}
