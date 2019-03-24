package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secret info!!")
}

// handleRequests - hendle all requests
func handleRequests() {

	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	fmt.Println("Hi server is here! - see me on port 8001")
	handleRequests()
}
