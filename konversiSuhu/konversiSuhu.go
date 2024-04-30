package main

import "fmt"

func main() {
	var celcius float32
	fmt.Println("Program Sederhana Konversi Suhu")
	fmt.Print("Masukan Suhu celcius :")
	fmt.Scan(&celcius)

	var reamur, fahrenheit, kelvin float32

	fmt.Println("Suhu Celcius : ", celcius)

	reamur = (4.0 / 5.0) * celcius
	fmt.Println("Suhu Reamur : ", reamur)

	fahrenheit = ((9.0 / 5.0) * celcius) + 32
	fmt.Println("Suhu Fahrenheit : ", fahrenheit)

	kelvin = celcius + 273
	fmt.Println("Suhu Kelvin : ", kelvin)
}
