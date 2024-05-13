package main

import "fmt"

func main() {
	angka := 10

	// Break
	for i := 1; i <= angka; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Angka ke:", i)
	}

	// Continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Bilangan Ganjil Ke :", i)
	}
}
