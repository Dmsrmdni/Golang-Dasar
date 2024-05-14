package main

import "fmt"

func main() {

	var cars map[string]string = map[string]string{}

	cars["color"] = "white"
	cars["merk"] = "bmw"

	fmt.Println(cars)      //map[color:white merk:bmw]
	fmt.Println(len(cars)) //2

	var biodata = map[string]string{
		"nama":     "Dimas",
		"addresse": "Bandung",
	}

	fmt.Println(biodata["nama"])
	fmt.Println(biodata["adress"])

	siswa := make(map[string]string)
	siswa["nama"] = "Dikha"
	siswa["jurusan"] = "Informatika"
	siswa["alamat"] = "Bandung"

	fmt.Println(siswa) //map[alamat:Bandung jurusan:Informatika nama:Dikha]

	delete(siswa, "alamat")

	fmt.Println(siswa) //map[jurusan:Informatika nama:Dikha]

}
