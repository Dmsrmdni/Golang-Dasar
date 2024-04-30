package main

import "fmt"

func main() {
	hari := "ss"

	switch hari {
	case "senin":
		fmt.Println("selamat Hari Senin")
	case "selasa":
		fmt.Println("selamat Hari selasa")
	case "rabu":
		fmt.Println("selamat Hari rabu")
	case "kamis":
		fmt.Println("selamat Hari kamis")
	case "jumat":
		fmt.Println("selamat Hari jumat")
	case "sabtu":
		fmt.Println("selamat Hari sabtu")
	case "minggu":
		fmt.Println("selamat Hari minggu")
	default:
		fmt.Println("tidak valid")
	}

	length := len(hari)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Terlalu Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
