package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	password := "mySecurePassword123"

	// Hash the password using MD5 (Insecure)
	hash := md5.Sum([]byte(password))

	fmt.Printf("MD5 Hash of password (Insecure): %x\n", hash)
}