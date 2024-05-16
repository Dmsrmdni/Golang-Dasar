package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func main() {

	var numberA int = 4

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}

	// di copy valuenya
	// address2 := address1
	// address2.Kota = "Sumedang"

	// fmt.Println(address1) //tidak berubah //{Bandung Jawa Barat Indonesia}

	// fmt.Println(address2) //{Sumedang Jawa Barat Indonesia}

	// langsung reference ke addres1
	address2 := &address1 //pointer
	address2.Kota = "Sumedang"

	fmt.Println(address1) //ikut berubah //{Sumedang Jawa Barat Indonesia}

	fmt.Println(address2) //{Sumedang Jawa Barat Indonesia}

}
