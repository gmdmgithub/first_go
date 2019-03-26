package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	validator "gopkg.in/go-playground/validator.v9"
)

// User - struct to keep data about user
type User struct {
	ID        string `json:"id" validate:"omitempty,uuid"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,gte=10"`
	Password  string `json:"password" validate:"required,gte=5"`
	Type      string `json:"type" validate:"required,gte=3"`
}

var mySignedKey = []byte("testmeifyoucanorjustuseown")

// GenerateJWT - func genereate JWT token
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claimes := token.Claims.(jwt.MapClaims)

	user := User{"bc4267fb-a38a-4b4a-9f98-172f0416bccb",
		"Alex", "Smith", "alex.smith@gmail.com", "alexsmith", "password", "admin"}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Validation passed!")
	}

	//validate part
	err = validate.StructExcept(user, "Password", "Type")
	if err != nil {
		log.Printf("Partial validation problem %v\t", err.Error())
	} else {
		log.Println("Partial validation passed!")
	}

	claimes["authorized"] = true
	claimes["client"] = "Alex Smith"
	claimes["user"] = user
	claimes["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySignedKey)

	if err != nil {
		err = fmt.Errorf("omething went wrong ... %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {

	log.Printf("Client got request")

	tokenValue, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("crash ...., %v\n", err))
	}
	respMap := make(map[string]string)
	respMap["token"] = tokenValue
	jsonString, err := json.Marshal(respMap)
	if err != nil {
		log.Println("cannot convert to JSON")
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	// fmt.Fprintf(w, string(jsonString))
	log.Println(string(jsonString))

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8001", nil)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	req.Header.Set("Token", tokenValue)

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	//finally replay
	fmt.Fprintf(w, string(body))

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
