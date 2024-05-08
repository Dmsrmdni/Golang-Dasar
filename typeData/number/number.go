package main

import "fmt"

func main() {
	// integer
	// var nilai8 int8 = 127
	// var nilai16 int16 = 32767
	// var nilai32 int32 = 2147483647
	// var nilai64 int64 = 9223372036854775807

	// fmt.Println(nilai8)  //127
	// fmt.Println(nilai16) //32767
	// fmt.Println(nilai32) //2147483647
	// fmt.Println(nilai64) //9223372036854775807

	// unsigned integer
	var nilai8 uint8 = 255
	var nilai16 uint16 = 65535
	var nilai32 uint32 = 4294967295
	var nilai64 uint64 = 18446744073709551615

	fmt.Println(nilai8)  //255
	fmt.Println(nilai16) //65535
	fmt.Println(nilai32) //4294967295
	fmt.Println(nilai64) //18446744073709551615

	nilai := 257
	var konversi = uint8(nilai)
	// jika melebihi batas maksimal akan kembali menjadi minimal
	fmt.Println(konversi) // 1
}
