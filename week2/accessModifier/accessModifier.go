package main

import (
	"belajar-golang-dasar-week2/accessModifier/biodata"
	"fmt"
)

func main() {
	fmt.Println("============Biodata============")
	// Dapat di akses dari biodata
	fmt.Println("Nama Saya : ", biodata.NamaLengkap)
	fmt.Println("Umur Saya : ", biodata.Umur)
	fmt.Println(biodata.SayHello("Dimas"))

	// tidak dapat di akses dari biodata
	// fmt.Println("Jurusan :", biodata.jurusan)
	// fmt.Println(biodata.sayGoodBye("Dimas"))

}
