package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	csvIFile  = "./sampleI.csv"
	csvOFile  = "./sampleO.csv"
	csvOMFile = "./sampleOM.csv"
)

func playCSV() {
	firstCSV()

	csvWithMap()

}

func csvWithMap() {

	us, err := readJSON(jsonFile)
	if err != nil {
		log.Printf("%s", err)
	}
	f, err := os.Create(csvOMFile)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	for i, u := range *us {

		var uM map[string]interface{}
		inVal, _ := json.Marshal(u)
		json.Unmarshal(inVal, &uM)

		// iterate through interfeces
		var titles []string
		var values []string
		for field, val := range uM {

			if i == 0 {
				titles = append(titles, field)
			}

			values = append(values, fmt.Sprintf("%v", val))
		}

		if len(titles) > 0 {
			w.Write(titles)
		}
		w.Write(values)
	}
	w.Flush()

	err = w.Error()
	if nil != err {
		log.Fatalln(err)
	}

}

func firstCSV() {

	us, err := readJSON(jsonFile)
	if err != nil {
		log.Printf("%s", err)
	}
	f, err := os.Create(csvOFile)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	// TODO! rething how to do it properly
	bytes, err := json.Marshal(us)
	if err != nil {
		log.Panicf("marshal problem %v", err)
	}
	// log.Printf("Bytes %v", bytes)

	s := make([]string, len(bytes))
	for i, val := range bytes {
		s[i] = string(val)
	}
	w.Write(s)

	w.Flush()

	err = w.Error()
	if nil != err {
		log.Fatalln(err)
	}
}
