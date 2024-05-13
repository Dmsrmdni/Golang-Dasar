package main

import "fmt"

func GetBiodata() (firstName, lastName, alamat string) {

	firstName = "Dimas"
	lastName = "Ramdani"

	return
}

func Calculate(x, y int) (penjumlahan, pengurangan int) {
	penjumlahan = x + y
	pengurangan = x - y
	return //secara otomatis akan mengembalikan penjumlahan dan pengurangan
}

func main() {
	firstName, lastName, alamat := GetBiodata()
	fmt.Println("Nama Saya", firstName, lastName)
	fmt.Println("Alamat Saya :", alamat)

	penjumlahan, pengurangan := Calculate(10, 5)
	fmt.Println("Total Penjumlahan :", penjumlahan)
	fmt.Println("Total Pengurangan :", pengurangan)

}
