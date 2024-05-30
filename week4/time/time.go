package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 time.Time = time.Now() //menggunakan waktu komputer kita
	fmt.Println(time1)               //2024-05-29 14:52:27.723337372 +0700 WIB m=+0.000041540
	fmt.Println(time1.Year())        //2024 tahun
	fmt.Println(time1.Month())       //May bulan
	fmt.Println(time1.Weekday())     //Wednesday hari
	fmt.Println(time1.Day())         //29 tanggal
	fmt.Println(time1.Hour())        //14 jam
	fmt.Println(time1.Minute())      //55 menit
	fmt.Println(time1.Second())      //33 detik

	var utc time.Time = time.Date(2010, time.January, 18, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local()) //konversi ke waktu lokal

	formatter := "2006-01-02 15:04:05"

	value := "2015-06-17 08:10:10"

	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}
