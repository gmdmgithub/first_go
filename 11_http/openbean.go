package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"log"
)

func openbeanRequest() {
	log.Println("Sending request to mockbin!")
	defer log.Println("End request to mockbin!")

	url := "http://mockbin.com/request?foo=bar&foo=baz"

	payload := strings.NewReader("{\"foo\": \"bar\"}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil{
		log.Printf("request problem %v",err)
		return
	}

	req.Header.Add("cookie", "foo=bar; bar=baz")
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-pretty-print", "4")

	res, err := http.DefaultClient.Do(req)
	if err !=nil{
		log.Printf("call problem %v", err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		log.Printf("read problem %v", err)
		return
	}

	fmt.Println(res)
	fmt.Println(string(body))

}