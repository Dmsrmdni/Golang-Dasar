package main

import "fmt"

func main() {
	var name string
	name = "Dimas Ramdani"
	fmt.Println(name)
	name = "Dikha adi nugraha"
	fmt.Println(name)

	// contoh penulisan lain
	// tanpa menuliskan type data apabila langsung di isi datanya
	var umur = 16
	fmt.Println(umur)

	// kata kunci var tidak wajib tetapi kita perlu menggunakan kata kunci := saat menganalisikan datanya (deklarasi pertama)
	hobby := "Menonton"
	fmt.Println("Hobby saya", hobby)

	// penulisan multipe variable
	var (
		firstName = "Dimas"
		lastName  = "Ramdani"
	)

	fmt.Println(firstName, lastName)

}
