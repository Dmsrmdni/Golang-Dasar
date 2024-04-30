package main

import "fmt"

func main() {
	for angka := 1; angka <= 10; angka++ {
		fmt.Printf("nilai sekarang adalah %v \n", angka)
	}

	haris := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	for index, hari := range haris {
		fmt.Println("index ke =", index, "hari ke =", hari)
	}

	// jika hanya ingin valuenya saja ganti variabel indexnya dengan _
	for _, hari := range haris {
		fmt.Println(hari)
	}
}
