package main

// Encrypting and decrypting data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"time"
)

// secretKey (key) parameter must be 16, 24 or 32, corresponding to the AES-128, AES-192 or AES-256 algorithms, respectively
// AES stands for Advanced Encryption Standard
var secretKey = []byte("this-is-the-most-secret-key-ever")

func encrypt(data, secretKey []byte) ([]byte, error) {

	// lets wrap it in Galois Counter Mode (GCM) with a standard nonce length
	gcm, err := getGCM(secretKey)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil

}

func decrypt(ciphertext []byte, secretKey []byte) ([]byte, error) {

	// lets wrap it in Galois Counter Mode (GCM) with a standard nonce length
	gcm, err := getGCM(secretKey)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

// getGCM - stands from Galois Counter Mode (GCM) - provides Authenticated Encryption with Associated Data (AEAD)
func getGCM(secretKey []byte) (cipher.AEAD, error) {

	//lets create a new block cipher based on the secretKey
	myCipher, err := aes.NewCipher(secretKey)
	//it could be also based on hashed secret key
	// myCipher,err := aes.NewCipher(createHash(passphrase)) //- from the previous example => 12_hashing
	if err != nil {
		return nil, err
	}
	// lets wrap it in Galois Counter Mode (GCM) with a standard nonce length
	gcm, err := cipher.NewGCM(myCipher)
	if err != nil {
		return nil, err
	}
	return gcm, nil
}

func main() {

	fmt.Println("welcome to the encrypt decrypt playground, the most opinioated AES (Advanced Encription Standard) is used")

	textToEnrypt := []byte("My name is Alex Smith - it's top secret information!")

	ciphertext, err := encrypt(textToEnrypt, secretKey)
	if err != nil {
		log.Fatal(err) //TODO - do something with this error
	}
	fmt.Printf("Planin text is: %s encrypted: %x\n", textToEnrypt, ciphertext)
	// ### decrypting
	plaintext, err := decrypt(ciphertext, secretKey)
	if err != nil {
		log.Fatal(err) // TODO: Properly handle error
	}
	fmt.Printf("Decrypted text: %x and encrypted is: %s\n", ciphertext, plaintext)

	//ltes write a file and decrypt it!
	encryptFile("test.txt", textToEnrypt, secretKey)
	time.Sleep(2 * time.Second)
	fmt.Printf("Secret data from file are: %s", decryptFile("test.txt", secretKey))
}
