package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

const (
	csvIFile = "./sampleI.csv"
	csvOFile = "./sampleO.csv"
)

func playCSV() {

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
