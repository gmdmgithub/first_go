package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getSampleData() {
	//####simple GET data
	log.Println("GET METHOD CALL")
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Printf("Something went wrong ...%v\n", err)
	} else {
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Body read problem ...%v\n", err)
		}
		fmt.Println(string(result))
	}
}

func postSampleData(url string) {

	log.Println("POST METHOD CALL")

	jsonSample := map[string]string{"name": "Alex", "email": "alex@gmail.com", "password": "123456"}

	jsonValue, err := json.Marshal(jsonSample)
	if err != nil {
		log.Printf("Post sample marchal problem %v\n", err)
		return
	}

	retriveResponse(http.Post(url, "application/json", bytes.NewBuffer(jsonValue)))

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

	retriveResponse(client.Do(request))

}

func retriveResponse(response *http.Response, err error) {

	if err != nil {
		log.Printf("Post sample post problem %v\n", err)
	} else {
		result, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			fmt.Println(string(result))
		}
	}
}

func main() {
	postURL := "https://httpbin.org/post"
	getSampleData()
	postSampleData(postURL)
	postWithHeader(postURL)

}
