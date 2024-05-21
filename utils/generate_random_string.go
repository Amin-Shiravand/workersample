package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// https://www.tutorialspoint.com/how-to-generate-random-string-characters-in-golang#:~:text=Generating%20random%20strings%20or%20characters%20in%20Golang%20can%20be%20achieved,random%20strings%20for%20cryptographic%20purposes.
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
