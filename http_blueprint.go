package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// define a type for the response
type Hello struct{}

// let that type implement the ServeHTTP method (defined in interface http.Handler)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there!")
}

func main() {
	var response Hello
	port := 4000
	fmt.Println("Server is running on port", port)
	http.ListenAndServe("localhost:"+strconv.Itoa(port), response)
}
