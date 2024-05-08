package main

import "fmt"

func main() {

	// fmt.Print("Dimas Ramdani") //Dimas Ramdani

	// Akan Menghasilkan baris baru
	fmt.Println("Dimas")
	fmt.Println("Ramdani")

	nama := "Dimas Ramdani"
	nilai := 80

	//Dengan format tertentu
	// mengambil type data nama -> string
	fmt.Printf("Type data nama : %T \n", nama)
	fmt.Printf("Nama saya : %s \n", nama) //Dimas Ramdani

	// mengambil type data nilai -> int
	fmt.Printf("Type data nilai : %T \n", nilai)
	fmt.Printf("Nama saya : %d \n", nilai) //80

	var umur int
	fmt.Print("masukan Umur :")
	fmt.Scan(&umur)
	fmt.Println("Umur anda", umur)
}
