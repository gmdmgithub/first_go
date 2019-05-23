package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func simpleJSON() {

	usr := User{
		Name:     "Alex",
		Password: "secret",
		Age:      27,
		Salary:   23400.00,
		Driver:   false,
	}

	// log.Printf("User is %v", usr)
	// var buf = new(bytes.Buffer)

	var b bytes.Buffer // A Buffer needs no initialization.

	usr.Write(&b)
	// log.Printf("Buffer is %s", b.String())

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
	// log.Printf("%s", buf)
}
