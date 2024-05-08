package main

import "fmt"

func main() {
	fmt.Println("Dimas")
	fmt.Println("Ramdani")
	var fullName string = "Dimas Ramdani"
	fmt.Println(fullName)

	// fungsi pada string
	/* menghitung jumlah panjang string
	menggunakan len('data')*/

	fmt.Println(len("Dimas"))

	/* mengambil karakter pada posisi yang di tentukan menggunakan
	[posisi karakternya contoh 0,1,2,3] hasilnya akan byte */
	fmt.Println("Dimas"[0])

}
