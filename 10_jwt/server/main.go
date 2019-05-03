package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("testmeifyoucanorjustuseown")

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secret info!!")
}

// isAuthorized - check if request is authorized
func isAuthorized(endpointToCheck func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0],
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error ")
					}
					return mySignedKey, nil
				})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				log.Printf("Token value %+v", token.Claims.(jwt.MapClaims)["client"])
				user := token.Claims.(jwt.MapClaims)["user"]
				log.Printf("User from token %+v", user)
				endpointToCheck(w, r)
			}

			return

		}

		fmt.Fprint(w, "Not authorized!")
	})

}

// handleRequests - handle all requests
func handleRequests() {

	// http.HandleFunc("/", mainPage)
	// after introduction middleware use isAuthorized

	http.Handle("/", isAuthorized(mainPage))

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	fmt.Println("Hi server is here! - see me on port 8001")
	handleRequests()
}
