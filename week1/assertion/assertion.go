package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nama interface{} = "Dimas Ramdani"
	resultString := nama.(string)
	fmt.Println(resultString) //Dimas Ramdani

	resultInt := nama.(int)
	fmt.Println(resultInt) //panic

	switch value := nama.(type) {
	case string:
		fmt.Println("tipe datanya adalah string", value)
	case int:
		fmt.Println("tipe datanya adalah integer", value)
	case bool:
		fmt.Println("tipe datanya adalah bool", value)
	default:
		fmt.Println("tipe datanya adalah", reflect.ValueOf(value))
	}
}
