package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Dimas"
	names[1] = "Ramdani"
	names[2] = "Dikha"

	fmt.Println(names)    // [Dimas Ramdani Dikha]
	fmt.Println(names[0]) // Dimas
	fmt.Println(names[2]) // Dikha

	var nilai = [...]int{1, 2, 3, 4, 5, 5, 6, 7, 1}
	fmt.Println(nilai)
	fmt.Println(len(nilai))

}
