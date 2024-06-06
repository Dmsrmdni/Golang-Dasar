package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	text1 := "Hello, world!"
	sha1 := sha1.New()
	sha1.Write([]byte(text1))
	encrypted1 := sha1.Sum(nil)
	encryptedString1 := fmt.Sprintf("%x", encrypted1)
	fmt.Println("Hash dari", text1, ":", encryptedString1)

	text2 := "Dimas Ramdani"
	sha1.Write([]byte(text2))
	encrypted2 := sha1.Sum(nil)
	encryptedString2 := fmt.Sprintf("%x", encrypted2)
	fmt.Println("Hash dari", text2, ":", encryptedString2)

}
