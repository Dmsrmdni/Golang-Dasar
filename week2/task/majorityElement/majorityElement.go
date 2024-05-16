package main

import "fmt"

func findMostFrequent(arr []int) int {
	// counts := make(map[int]int)
	var counts map[int]int = map[int]int{}

	// Menghitung kemunculan setiap nilai dalam array
	for _, num := range arr {
		counts[num]++
	}

	// Variabel untuk menyimpan nilai terbanyak dan jumlah kemunculannya
	nilaiTerbanyak := 0
	jumlahMuncul := 0

	// Mencari nilai terbanyak
	for num, count := range counts {
		if count > jumlahMuncul { //2 > 0 = true ||  1 > 2 = false
			nilaiTerbanyak = num
			jumlahMuncul = count // 2
		}
	}

	return nilaiTerbanyak
}

func main() {
	nilai := []int{3, 2, 1}

	fmt.Println(findMostFrequent(nilai))
}
