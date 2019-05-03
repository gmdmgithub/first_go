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
	Username  string `json:"username" validate:"required,gte=6"`
	Password  string `json:"password" validate:"required,gte=5"`
	Type      string `json:"type" validate:"required,gte=3"`
}

var mySignedKey = []byte("testmeifyoucanorjustuseown")

// GenerateJWT - func genereate JWT token
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	user := User{"bc4267fb-a38a-4b4a-9f98-172f0416bccb",
		"Alex", "Smith", "alex.smith@gmail.com", "alexsmith", "password", "admin"}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	log.Println("Validation passed!")

	//validate part
	err = validate.StructExcept(user, "Password", "Type")
	if err != nil {
		log.Printf("Partial validation problem %v\t", err.Error())
		return "", err
	}

	log.Println("Partial validation passed!")

	claims["authorized"] = true
	claims["client"] = "Alex Smith"
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySignedKey)

	if err != nil {
		err = fmt.Errorf("something went wrong ... %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {

	log.Printf("Client got request")

	tokenValue, err := GenerateJWT()
	if err != nil {
		log.Printf("GenerateJWT problem %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Print("Token generated")
	respMap := make(map[string]string)
	respMap["token"] = tokenValue

	if err := json.NewEncoder(w).Encode(respMap); err != nil {
		log.Printf("Encode problem %v", err)
	}

	log.Printf("json string passed")

	client := &http.Client{}

	//send request to the server
	req, err := http.NewRequest("GET", "http://localhost:8001", nil)
	if err != nil {
		log.Printf("NewRequest problem %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Token", tokenValue)
	// final sanding the request
	res, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do(req) problem %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ReadAll problem %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
