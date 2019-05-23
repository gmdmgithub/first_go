package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
)

const (
	jsonFile = "./users.json"
	xmlFile  = "./users.xml"
)

func playXML() {
	us, err := readJSON(jsonFile)
	if err != nil {
		log.Printf("%s", err)
	}

	f, err := os.Create(xmlFile)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	if err := xml.NewEncoder(f).Encode(us); err != nil {
		log.Printf("xml encode problem %v", err)
	}

}

func readJSON(name string) (u *[]User, err error) {

	f, err := os.Open(name)
	if err != nil {
		log.Printf("file read problem %v", err)
		return nil, err
	}
	defer f.Close()

	var usrL Users

	if err := json.NewDecoder(f).Decode(&usrL); err != nil {
		log.Printf("decode problem %v", err)
		return nil, err
	}

	// log.Printf("%v", usrL)

	return &usrL.UserL, nil

}
