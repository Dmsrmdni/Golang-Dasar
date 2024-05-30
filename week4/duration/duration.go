package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 6 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond

	fmt.Println(duration1) //6s
	fmt.Println(duration2) //10ms

	time.Sleep(time.Second * 3)
	fmt.Println("haii!!!")
}
