package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	rand.Seed(time.Now().UnixNano()) // (secure)

	token := fmt.Sprintf("%x", rand.Int63()) // (insecure)

	fmt.Println("Generated Token (Insecure):", token)
}