package main

import "fmt"

func print(value interface{}) any {
	return value
}
func main() {
	fmt.Println(print("hello world"))
	fmt.Println(print(123))
	fmt.Println(print(true))
}
