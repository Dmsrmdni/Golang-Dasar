package main

import "fmt"

func main() {
	angka := 10

	for i := 1; i <= angka; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Angka ke:", i)
	}
}
