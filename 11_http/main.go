package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	// testURL - url  from API whitch simple retern what get
	testURL = "https://httpbin.org/[REPLACER]"
	// repl - string to replace
	repl = "[REPLACER]"
)

func getSampleData() {
	//####simple GET data
	log.Println("GET METHOD CALL")
	url := strings.Replace(testURL, repl, "get", 1)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Something went wrong ...%v\n", err)
		return
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Body read problem ...%v\n", err)
		return
	}
	fmt.Println(string(result))

}

func postSampleData(url string) {

	log.Println("POST METHOD CALL")

	jsonSample := map[string]string{"name": "Alex", "email": "alex@gmail.com", "password": "123456"}

	jsonValue, err := json.Marshal(jsonSample)
	if err != nil {
		log.Printf("Post sample marchal problem %v\n", err)
		return
	}

	resp, err := retriveResponse(http.Post(url, "application/json", bytes.NewBuffer(jsonValue)))
	if err != nil {
		log.Println("problem with retriving data", err)
		return
	}
	fmt.Println("postSampleData", resp)
}

func postWithHeader(url string) {
	log.Println("POST WITH HEADER METHOD CALL")

	jsonSample := map[string]string{"name": "Max", "email": "max@gmail.com", "password": "test"}

	jsonValue, err := json.Marshal(jsonSample)
	if err != nil {
		log.Printf("Post sample marchal problem %v\n", err)
		return
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("request problem %v\n", err)
		return
	}
	//we can add header
	request.Header.Set("Content-Type", "application/json")

	// now we can create client request
	client := &http.Client{}

	resp, err := retriveResponse(client.Do(request))
	if err != nil {
		log.Println("postWithHeader problem", err)
	}
	fmt.Println("postWithHeader result: ", resp)

}

func retriveResponse(response *http.Response, err error) (string, error) {

	if err != nil {
		log.Printf("Post sample post problem %v\n", err)
		return "", err
	}
	result, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return "", err
	}
	// fmt.Println(string(result))
	return string(result), nil

}

func main() {
	postURL := strings.Replace(testURL, repl, "post", 1)
	getSampleData()
	postSampleData(postURL)
	postWithHeader(postURL)

	openbeanRequest()
}
