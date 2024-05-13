package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var player, computer, hasil string

	randomNumber := rand.Intn(3) + 1

	if randomNumber == 1 {
		computer = "batu"
	} else if randomNumber == 2 {
		computer = "gunting"
	} else if randomNumber == 3 {
		computer = "kertas"
	}

	fmt.Println("Pilihlah antara batu, gunting, kertas")
	fmt.Print("Masukan pilihan Player :")
	fmt.Scan(&player)

	if player == computer {
		hasil = "seri"
	} else if player == "batu" {
		if computer == "kertas" {
			hasil = "kalah"
		} else {
			hasil = "menang"
		}
	} else if player == "gunting" {
		if computer == "kertas" {
			hasil = "menang"
		} else {
			hasil = "kalah"
		}
	} else if player == "kertas" {
		if computer == "gunting" {
			hasil = "kalah"
		} else {
			hasil = "menang"
		}
	} else {
		hasil = "tidak valid"
	}
	fmt.Println("Player memilih :", player, "Computer memilih :", computer)
	fmt.Println("hasilnya :", hasil)
}
