package main

import "fmt"

func main() {
	// Deklarasi variabel beserta type data
	var myName string = "Dimas Ramdani"
	fmt.Println(myName) //Dimas Ramdani

	// tanpa menuliskan type data apabila langsung di isi datanya
	var umur = 19
	fmt.Println(umur)

	/* kata kunci var tidak wajib tetapi kita perlu menggunakan
	kata kunci := saat menganalisikan datanya (deklarasi pertama) */
	hobby := "Menonton"
	fmt.Println("Hobby saya", hobby)

	// penulisan multiple variable
	var (
		firstName = "Dimas"
		lastName  = "Ramdani"
	)

	fmt.Println(firstName, lastName)

	var first, second, third string = "satu", "dua", "tiga"
	fmt.Println(first, second, third) // satu dua tiga
	satu, dua, tiga := 1, "dua", 3
	fmt.Println(satu, dua, tiga) // 1 dua 3

	// var title string // Error karena tidak di gunakan

	/* Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung
	nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah. */
	name, _ := "Dimas Ramdani", 2004

	fmt.Println(name)
}
