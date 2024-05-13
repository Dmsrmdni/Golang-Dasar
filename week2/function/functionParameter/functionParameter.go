package main

import "fmt"

func LuasPersegiPanjang(panjang int, lebar int) {
	luas := panjang * lebar

	fmt.Println("Panjang :", panjang, "Lebar :", lebar)
	fmt.Println("Luas Persegi panjangnya adalah :", luas)
}

func main() {
	LuasPersegiPanjang(15, 5)
	LuasPersegiPanjang(10, 4)
}
