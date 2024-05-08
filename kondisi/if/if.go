package main

import "fmt"

func main() {
	rataRata := 85
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

	fmt.Println(grade) //Grade B

	// short statement
	if nilai := 90; nilai > 70 {
		fmt.Println("Kamu Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
