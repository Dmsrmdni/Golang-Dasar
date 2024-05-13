package main

import "fmt"

func main() {
	hari := "senin"

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

	nilai := 80
	// kondisi di switch tidak wajib, bisa menambahkan pada case nya
	switch {
	case nilai >= 80:
		fmt.Println("Grade A")
	case nilai >= 70:
		fmt.Println("Grade B")
	default:
		fmt.Println("Grade C")
	}
}
