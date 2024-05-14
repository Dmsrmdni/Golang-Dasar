package main

import "fmt"

type Siswa struct {
	Nama, Alamat string
	Umur         int
}

func (siswa Siswa) SayHello(name string) {
	fmt.Println("Hello", name, "Nama saya", siswa.Nama)
}

func main() {
	var dimas Siswa

	dimas.Nama = "Dimas Ramdani"
	dimas.Alamat = "Bandung"
	dimas.Umur = 19

	fmt.Println(dimas)      //{Dimas Ramdani Bandung 19}
	fmt.Println(dimas.Nama) //Dimas Ramdani

	dikha := Siswa{
		Nama:   "Dikha",
		Alamat: "Sumedang",
		Umur:   20,
	}

	// SayHello() //error tidak bisa langsung di akses
	dikha.SayHello("dimas")

	fmt.Println(dikha) //{Dikha Sumedang 20}
}
