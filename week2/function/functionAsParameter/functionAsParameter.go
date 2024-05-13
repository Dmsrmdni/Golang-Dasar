package main

import "fmt"

func grad(nama string, nilai int, filter func(int) string) {
	gradeFilter := filter(nilai)

	fmt.Printf("Nama : %v, Nilai : %v, Grade : %v\n", nama, nilai, gradeFilter)
}

func gradeFilter(nilai int) string {
	if nilai >= 85 {
		return "Grade A"
	} else if nilai >= 75 {
		return "Grade B"
	} else {
		return "Grade C"
	}
}

func main() {

	grad("Dimas", 90, gradeFilter)
	grad("Ramdani", 70, gradeFilter)
}
