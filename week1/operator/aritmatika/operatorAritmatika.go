package main

import "fmt"

func main() {
	angka1 := 10
	angka2 := 5

	fmt.Printf("angka 1 = %v dan angka 2 = %v\n", angka1, angka2)
	fmt.Println("penjumlahan")
	penjumlahan := angka1 + angka2
	fmt.Println("Hasil Penjumlahan =", penjumlahan)

	fmt.Println("pengurangan")
	pengurangan := angka1 - angka2
	fmt.Println("Hasil Pengurangan =", pengurangan)

	fmt.Println("perkalian")
	perkalian := angka1 * angka2
	fmt.Println("Hasil Perkalian =", perkalian)

	fmt.Println("Pembagian")
	pembagian := angka1 / angka2
	fmt.Println("Hasil Pembagian =", pembagian)

	fmt.Println("Sisa Bagi")
	modulus := (angka1 % 2)
	fmt.Println("Hasil sisa pembagian 2 =", modulus)

}
