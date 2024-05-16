package main

import "fmt"

func selesai() {
	fmt.Println("Terimakasih")

	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func mulai() {
	/* akan di eksekusi di akhir
	dan apabila ada error juga akan tetap di eksekusi */

	defer selesai()
	if error := true; error {
		// menghetikan eksekusi program
		panic("Error!!!!!")
	}

	/*jika penulisan seperti ini tidak akan tereksekusi
	jadi harus di simpan pada defer */

	// message := recover()
	// fmt.Println("Terjadi Panic", message)
}

func main() {
	mulai()
	fmt.Println("Dimas Ramdani")
}
