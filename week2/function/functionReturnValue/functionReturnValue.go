package main

import "fmt"

func LuasPersegi(sisi int) int {
	return sisi * sisi
}

func main() {
	luas := LuasPersegi(10)
	fmt.Println(luas)

	// secara langsung tanpa menggunakan variabel
	fmt.Println(LuasPersegi(5))
}
