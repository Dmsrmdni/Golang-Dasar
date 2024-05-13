package main

import "fmt"

func main() {
	angka1 := 5
	angka2 := 5

	lebihDari := angka1 > angka2
	lebihDariSamaDengan := angka1 >= angka2

	kurangDari := angka1 < angka2
	kurangDariSamaDengan := angka1 <= angka2

	samaDengan := angka1 == angka2

	tidakSamaDengan := angka1 != angka2

	fmt.Println("Lebih dari :", lebihDari)                       //false
	fmt.Println("Lebih dari sama dengan :", lebihDariSamaDengan) //true

	fmt.Println("Kurang dari :", kurangDari)                       //false
	fmt.Println("Kurang dari sama dengan :", kurangDariSamaDengan) //true

	fmt.Println("Sama dengan :", samaDengan) //true

	fmt.Println("Tidak Sama dengan :", tidakSamaDengan) //false

}
