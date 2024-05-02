package main

import (
	"fmt"
	"strconv"
)

func main() {
	// - Int32 to Int64
	var numb32 int32 = 123
	numb32ToNumb64 := int64(numb32)
	fmt.Println("Konversi int 32 to int64 dari nilai :", numb32, "Adalah", numb32ToNumb64)

	// - Int64 to Int32
	var numb64 int64 = 700
	numb64ToNumb32 := int32(numb64)
	fmt.Println("Konversi int 32 to int64 dari nilai :", numb64, "Adalah", numb64ToNumb32)

	// - String to Byte
	var str1 string = "Dimas"
	stringToByte := []byte(str1)
	fmt.Println("Konversi string to byte dari :", str1, "Adalah", stringToByte)

	// - String to Rune
	var str2 string = "Ramdani"
	stringToRune := []rune(str2)
	fmt.Println("Konversi string to rune dari :", str2, "Adalah", stringToRune)

	// - Int64 to String
	var numb64_2 int64 = 300
	numb64ToString := strconv.FormatInt(numb64_2, 10)
	fmt.Println("Konversi int64 to string dari :", numb64_2, "Adalah", numb64ToString)

	// - Int32 to String
	var numb32_2 int32 = 2147483647
	numb32ToString := strconv.Itoa(int(numb32_2))
	fmt.Println("Konversi int64 to string dari :", numb32_2, "Adalah", numb32ToString)

	// - String to Int64
	var str3 string = "324"
	stringToInt64, _ := strconv.ParseInt(str3, 10, 64)
	fmt.Println("Konversi string to int64 dari :", str3, "Adalah", stringToInt64)

	// - String to Int32
	var str4 string = "32422"
	stringToInt32, _ := strconv.ParseInt(str4, 10, 32)
	fmt.Println("Konversi string to int32 dari :", str4, "Adalah", stringToInt32)

	// - String to Bool
	var str5 = "true"
	var bul, err = strconv.ParseBool(str5)
	if err == nil {
		fmt.Println(bul)
	}

}
