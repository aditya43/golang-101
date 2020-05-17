package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("Aditya")
	inputPassword := []byte("Aditya1")

	hash := generateHash(password)
	compareHash(hash, inputPassword)
}

func generateHash(password []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(password, 4)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bcrypt Hash: %v", string(hash))

	return hash
}

// CompareHashAndPassword returns 'nil' on success and error on failure
func compareHash(hash []byte, inputPassword []byte) {
	err := bcrypt.CompareHashAndPassword(hash, inputPassword)

	if err != nil {
		// Password doesn't match
		log.Fatal(err)
	}
}
