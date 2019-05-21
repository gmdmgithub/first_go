package main

import (
	"encoding/json"
	"io"
	"log"
)

type User struct { //TODO!! - add more and remember Capitalize first letter!!
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
	Salary   float32 `json:"salary"`
	Driver   bool    `json:"driver"`
}

type Users struct {
	UserL []User `json:"users"`
}

// The method takes just a io.Writer as input
func (p *User) Write(w io.Writer) {
	b, err := json.Marshal(*p)
	if err != nil {
		log.Printf("problem with marchaling %v", err)
	}
	// Inside our function we just write into the io.Writer
	// We don't care about which writer we use
	w.Write(b)
}
