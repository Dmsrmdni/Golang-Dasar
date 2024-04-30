package main

import "fmt"

func main() {
	// var i int

	// fmt.Print("masukan Nomber 1-10 :")
	// fmt.Scan(&i)
	// fmt.Println("Nomor Yang Anda Masukan", i)

	// var i, j int

	// fmt.Print("masukan Nomber 1-10 :")
	// fmt.Scanln(&i, &j)
	// fmt.Println("Nomor Yang Anda Masukan", i, "and", j)

	var nama, kelas string

	fmt.Print("Masukan Nama :")
	fmt.Scanln(&nama)

	fmt.Print("Masukan Kelas :")
	fmt.Scanln(&kelas)

	fmt.Println("Nama Saya", nama, "kelas", kelas)
}
