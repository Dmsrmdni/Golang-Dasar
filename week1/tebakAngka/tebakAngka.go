package main

import (
	"fmt"
	"math/rand"
)

func main() {
	kesempatan := 3
	var tebakan int

	randomNumber := rand.Intn(10 - 1)
	fmt.Println("Selamat Bermain")

	for kesempatan > 0 {
		fmt.Printf("anda Memiliki %v Kesempatan \n", kesempatan)
		fmt.Println("======================")
		fmt.Print("Masukan Tebakan anda (1-10) : ")
		fmt.Scan(&tebakan)
		if randomNumber == tebakan {
			fmt.Println("Anda Menang Angka yang di cari adalah", randomNumber)
			break
		} else if randomNumber > tebakan {
			kesempatan--
			fmt.Println("======================")
			fmt.Println("Angka Terlalu Rendah")
		} else if randomNumber < tebakan {
			kesempatan--
			fmt.Println("======================")
			fmt.Println("Angka Terlalu tinggi")
		}
	}

	if kesempatan == 0 {
		fmt.Println("Kamu Kalah kesempatan Habis Angka yang dicari adalah", randomNumber)
	}

}
