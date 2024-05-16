package main

import (
	"errors"
	"fmt"
)

func perkalian(nilai, kali int) (int, error) {
	if kali == 0 {
		return 0, errors.New("perkalian dengan nol")
	} else {
		return nilai * kali, nil
	}
}

func main() {
	hasil, err := perkalian(10, 0)

	if err == nil {
		fmt.Println("Hasilnya : ", hasil)
	} else {
		fmt.Println("Error : ", err)
	}
}
