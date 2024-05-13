package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {
	fmt.Println(sum(10, 20, 2, 3, 4, 2, 4, 9)) //54
	fmt.Println(sum(10, 20, 2))                //32

	nums := []int{1, 2, 3, 4, 5}

	// untuk memecah slice menggunakan ...
	fmt.Println(sum(nums...)) //15

}
