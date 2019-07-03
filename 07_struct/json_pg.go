package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Author - sample stract representing Author of the book
type Author struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Known bool   `json:"is_known"`
}

// Book - sample struct representing book
type Book struct {
	Title  string `json:"title,omitempty"`
	Author Author `json:"author,omitempty"`
}

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

	book := Book{Title: "Go best programming language!",
		Author: Author{ID: 1, Name: "Rob Pike", Age: 24, Known: false}}
	log.Printf("My book is %+v", book)

	// byteArray, err := json.Marshal(book)
	byteArray, err := json.MarshalIndent(book, "", "	") //with intend
	if err != nil {
		log.Printf("Marshal problem %s", err)
		return
	}

	log.Printf("Stringed marshaled book is %v", string(byteArray))

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

// Truck - sample struct
type Truck struct {
	Name     string `json:"name,omitempty"`
	Make     string `json:"make,omitempty"`
	MaxSpeed int    `json:"max_speed,omitempty"`
}

func streamJSONUnmarshal() {
	ins := `[{"Name":"X3","Make":"BMW","MaxSpeed":300},{"Name":"A6","Make":"Audi","MaxSpeed":330}]`

	in := `{"Name":"Golf","Make":"WW","MaxSpeed":220}`

	trs := []Truck{}
	err := json.Unmarshal([]byte(ins), &trs)
	if err != nil {
		log.Printf("Unmarshal problem %s", err)
		return
	}
	fmt.Printf("Results are: %+v\n", trs)

	tr := Truck{}
	err = json.Unmarshal([]byte(in), &tr)
	if err != nil {
		log.Printf("Unmarshal problem %s", err)
		return
	}
	fmt.Printf("Result is: %+v\n", tr)

	// IMPORTANT _ DIFFERENT APPROACH!!!
	var mapInterface map[string]interface{}
	err = json.Unmarshal([]byte(in), &mapInterface)
	if err != nil {
		log.Printf("Unmarshal problem %s", err)
		return
	}
	fmt.Printf("Result with map is: %+v\n", mapInterface)
}
