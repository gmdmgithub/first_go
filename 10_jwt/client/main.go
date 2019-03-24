package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("testmeifyoucanorjustuseown")

// GenerateJWT - func genereate JWT token
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claimes := token.Claims.(jwt.MapClaims)

	claimes["authorized"] = true
	claimes["client"] = "Alex"
	claimes["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySignedKey)

	if err != nil {
		err = fmt.Errorf("Something went wrong ... %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {

	log.Printf("Client get request /")

	tokenValue, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("crash ...., %v\n", err))
	}
	fmt.Fprintf(w, "got token %s", tokenValue)
}

func handleRequest() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("Hi there client here, use me on port 8081")

	// tokenValue, err := GenerateJWT()
	// if err != nil {
	// 	fmt.Printf("crash ...., %v\n", err)
	// }
	// fmt.Printf("got token %s", tokenValue)

	handleRequest()

}
