package main

import "fmt"

func main() {
	var myName string = "Dimas Ramdani"
	fmt.Println(myName)

	var yourName string
	yourName = "Dikha adi nugraha"
	fmt.Println(yourName)

	// contoh penulisan lain
	// tanpa menuliskan type data apabila langsung di isi datanya
	var umur = 16
	fmt.Println(umur)

	// kata kunci var tidak wajib tetapi kita perlu menggunakan kata kunci := saat menganalisikan datanya (deklarasi pertama)
	hobby := "Menonton"
	fmt.Println("Hobby saya", hobby)

	// penulisan multiple variable
	var (
		firstName = "Dimas"
		lastName  = "Ramdani"
	)

	fmt.Println(firstName, lastName)

}
