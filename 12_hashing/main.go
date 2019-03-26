package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	bcrypt "golang.org/x/crypto/bcrypt"
)

// createHash - one way hash method - not recomended for securing
func createHash(key string) string {
	//sha 256
	hash1 := sha256.New()
	io.WriteString(hash1, key)
	fmt.Printf("%x\n", hash1.Sum(nil))

	hash2 := sha1.New()
	io.WriteString(hash2, key)
	fmt.Printf("%x\n", hash2.Sum(nil))

	//md5 hash
	hashMd5 := md5.New()

	io.WriteString(hashMd5, key)
	fmt.Printf("%x\n", hashMd5.Sum(nil))

	//with Write method
	hashMd5_2 := md5.New()
	hashMd5_2.Write([]byte(key))
	return hex.EncodeToString(hashMd5_2.Sum(nil))
}

func passwordEncrypter(password string) []byte {

	encrptedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(encrptedPass))
	// TODO: Store this pass in DB

	return encrptedPass
}

func checkPassword(passwod, hashedPassword []byte) bool {

	if err := bcrypt.CompareHashAndPassword(hashedPassword, passwod); err != nil {
		// TODO: Properly handle error
		log.Println(err)
		return false
	}
	// TODO: login user or something
	log.Println("Password is corrct!!!")
	return true
}

func main() {

	fmt.Println("Here you have one way hashing - very useful for password (one way encryption!)")

	fmt.Printf(createHash("Secret information"))
	mySecrretPassword := "123456"
	pass := passwordEncrypter(mySecrretPassword)
	fmt.Printf("My secret password is : %s", pass)

	checkPassword([]byte(mySecrretPassword), pass)
	checkPassword([]byte("alex"), pass)

}
