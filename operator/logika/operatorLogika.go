package main

import "fmt"

func main() {
	mapel1 := 80
	mapel2 := 90

	lulusMapel1 := mapel1 > 80
	lulusMapel2 := mapel2 > 70

	dan := lulusMapel1 && lulusMapel2
	atau := lulusMapel1 || lulusMapel2

	fmt.Println(dan)
	fmt.Println(!dan)
	fmt.Println(atau)
	fmt.Println(!atau)

}
