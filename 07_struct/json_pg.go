package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type User struct { //TODO!! - add more and remember Capitalize first letter!!
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int
}

type Users []User

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

func simpleJSON() {

	usr := User{
		Name:     "Alex",
		Password: "secret",
		Age:      27,
	}

	log.Printf("User is %v", usr)
	// var buf = new(bytes.Buffer)

	var b bytes.Buffer // A Buffer needs no initialization.

	usr.Write(&b)
	log.Printf("Buffer is %s", b.String())

	// io.Copy(os.Stdout, buf)

}

func withBuffer() {
	log.Println("Buffer and Encode way")
	usr := User{
		Name:     "Kris",
		Password: "1111",
		Age:      31,
	}
	var buf = new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(usr); err != nil {
		log.Printf("Problem with decode %v", err)
	}
	log.Printf("%s", buf)
}
