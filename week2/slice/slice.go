package main

import "fmt"

func main() {
	var days = []string{"senin", "selasa", "rabu", "kamis", "jumat"} //slice
	var fruits = [4]string{"apple", "grape", "banana", "melon"}      //array

	fmt.Println(days)
	fmt.Println(fruits)

	daysNew := append(days, "sabtu")
	fmt.Println(daysNew) //[senin selasa rabu kamis jumat sabtu]

	// Membuat slice dari array dimulai index low sampai index sebelum high
	fruits1 := fruits[0:2]
	fmt.Println(fruits1) //[apple grape]

	// Membuat slide dari array dimulai index low sampai index akhir di array
	fruits2 := fruits[2:]
	fmt.Println(fruits2) //[banana melon]

	// Membuat slice dari array dimulai index 0 sampai index sebelum high
	fruits3 := fruits[:3]
	fmt.Println(fruits3) //[apple grape banana]

	// Membuat slice dari array dimulai index 0 sampai index akhir di array
	fruits4 := fruits[:]
	fmt.Println(fruits4)      //[apple grape banana melon]
	fmt.Println(len(fruits4)) //4
	fmt.Println(cap(fruits4)) //4
}
