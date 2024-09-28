package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	text := "Unsafe MD5 Hash"

	hash := md5.New()
	hash.Write([]byte(text))

	checksum := hex.EncodeToString(hash.Sum(nil))

	fmt.Printf("MD5 hash of '%s' is: %s\n", text, checksum)
}s