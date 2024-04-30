package main

import "fmt"

func main() {
	rataRata := 50
	var grade string

	if rataRata >= 90 {
		grade = "Grade A"
	} else if rataRata >= 80 {
		grade = "Grade B"
	} else if rataRata >= 70 {
		grade = "Grade C"
	} else {
		grade = "Grade D"
	}

	fmt.Println(grade)

	if length := len("Dimas Ramdani"); length > 5 {
		fmt.Println("sudah mantap")
	} else {
		fmt.Println("nama anda terlalu pendek")
	}
}
