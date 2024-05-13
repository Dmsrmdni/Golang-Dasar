package main

import "fmt"

func getBiodata(namaLengkap string, umur int, alamat string) (string, int, string) {
	return namaLengkap, umur, alamat
}

func main() {
	nama, umur, alamat := getBiodata("Dimas Ramdani", 19, "Bandung")
	fmt.Println("Nama Saya :", nama)
	fmt.Println("Umur Saya :", umur)
	fmt.Println("Alamat Saya :", alamat)

	// Menghiraukan return value menggunakan (_)
	nama2, _, _ := getBiodata("Dikha", 19, "Sumedang")
	fmt.Println("Nama :", nama2)

}
