package main

import "fmt"

func main() {
	const firstName string = "Dimas"
	const lastName = "Ramdani"

	//error jika ingin mengubah constant
	// firstName = "error ubah data const"

	fmt.Println(firstName, lastName)

	// Penulisan Multiple constanta
	const (
		umur  = 19
		hobby = "Menonton"
	)

	fmt.Println("Umur Saya :", umur, "Hobby saya :", hobby)

	const satu, dua = 1, 2

	fmt.Println(satu, dua)
}
