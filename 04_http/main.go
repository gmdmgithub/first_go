package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// http.ResponseWriter assembles the servers response and writes to
// the client
// http.Request is the clients request
func handler(w http.ResponseWriter, r *http.Request) {

	// Writes to the client
	fmt.Fprintf(w, "Hi there we are with you and using Go!\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	users := []string{"Mark", "Frank", "John"}
	fmt.Fprintf(w, "Hello on Earth\n")
	for i := 0; i < len(users); i++ {
		fmt.Fprintln(w, users[i])
	}
}

func handlerOnce() http.HandlerFunc {
	log.Println("This code will execute only once!")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("This code is like clousers")
		fmt.Fprintf(w, "Hi there - back from func")
	}
}

func main() {

	// Calls for function handlers output to match the directory /
	http.HandleFunc("/", handler)

	// Calls for function handler2 output to match directory /users
	http.HandleFunc("/users", handler2)

	// Calls for function handler2 output to match directory /users
	http.HandleFunc("/hello", handlerOnce())

	port := 8081
	// Listen to port (8081) and handle requests - response is nil (null)
	fmt.Println("Server is running on port " + strconv.Itoa(port) + " ...")
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
