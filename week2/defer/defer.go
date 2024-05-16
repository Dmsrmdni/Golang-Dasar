package main

import "fmt"

func selesai() {
	fmt.Println("Terimakasih")
}

func mulai() {
	/* akan di eksekusi di akhir
	dan apabila ada error juga akan tetap di eksekusi */
	defer selesai()
	fmt.Println("Selamat Datang")

	// selesai() //jika terjadi error tidak akan di eksekusi
}

func main() {
	mulai()
}
