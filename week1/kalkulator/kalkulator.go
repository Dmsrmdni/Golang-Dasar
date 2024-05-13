package main

import "fmt"

func main() {
	var pilihan int
	fmt.Println("Pilih salah satu")
	fmt.Println("1. Penjumlahan (+)")
	fmt.Println("2. Pengurangan (-)")
	fmt.Println("3. Perkalian (x)")
	fmt.Println("4. Pembagian (/)")

	fmt.Print("Pilih antara (1 - 4) :")
	fmt.Scan(&pilihan)

	if pilihan == 1 || pilihan == 2 || pilihan == 3 || pilihan == 4 {
		var angka1, angka2, hasil int
		fmt.Print("Masukan Angka Pertama :")
		fmt.Scan(&angka1)

		fmt.Print("Masukan Angka Kedua :")
		fmt.Scan(&angka2)

		if pilihan == 1 {
			hasil = (angka1 + angka2)
		} else if pilihan == 2 {
			hasil = (angka1 - angka2)
		} else if pilihan == 3 {
			hasil = (angka1 * angka2)
		} else if pilihan == 4 {
			hasil = (angka1 / angka2)
		}

		fmt.Println("Hasilnya adalah", hasil)

	} else {
		fmt.Println("Pilihan yang anda masukan tidak valid")
	}

}
