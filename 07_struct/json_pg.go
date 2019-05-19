package main

import (
	"bytes"
	"encoding/json"
	"log"
)

type User struct {
	name     string
	password string
	age      int
	salary   float32
	manager  bool
}

type Users []User

func simpleJSON() {

	usr := User{
		name:     "Alex",
		password: "secret",
		age:      27,
		salary:   120000.00,
		manager:  false,
	}

	log.Printf("User is v%", usr)

	var buf = new(bytes.Buffer)

	err := json.NewEncoder(buf).Encode(usr)
	if err != nil {
		log.Fatal("problem with marchaling ", err)
	}

	log.Printf("Encodeing result %v", buf)

	// io.Copy(os.Stdout, buf)

}
