package main

import (
	"io/ioutil"
	"os"
)

func encryptFile(filename string, data, secretKey []byte) {
	f, _ := os.Create(filename)
	defer f.Close()
	encryptedData, err := encrypt(data, secretKey)
	if err != nil {
		return
	}
	f.Write(encryptedData)
}

func decryptFile(filename string, secretKey []byte) []byte {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	decryptedData, err := decrypt(data, secretKey)
	if err != nil {
		return nil
	}
	return decryptedData
}
