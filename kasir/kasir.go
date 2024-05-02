package main

import "fmt"

func main() {
	var mauMakan, mauMinuman, pilihMinuman, pilihMakanan, namaMinuman, namaMakanan string
	hargaMakanan, hargaMinuman, totalHarga, pembayaran, kembalian, jumlahMakanan, jumlahMinuman := 0, 0, 0, 0, 0, 0, 0

	fmt.Println("========================================")
	fmt.Println("|| Selamat datang di rumah makan kami ||")
	fmt.Println("========================================")
	fmt.Print("Ingin memesan makanan ? (iya/tidak) :")
	fmt.Scan(&mauMakan)

	if mauMakan == "iya" {
		fmt.Println("Menu makanan :")
		fmt.Println("1. Nasi Goreng | Rp : 15000")
		fmt.Println("2. Mie Goreng  | Rp : 10000")
		fmt.Println("3. Ayam Katsu  | Rp : 20000")
		fmt.Println("4. Sate        | Rp : 15000")
		fmt.Println("5. Ayam Balado | Rp : 17000")

		fmt.Print("Pilih Menu (1 - 5) :")
		fmt.Scan(&pilihMakanan)

		fmt.Print("jumlah Pesanan :")
		fmt.Scan(&jumlahMakanan)

		if pilihMakanan == "1" {
			namaMakanan = "Nasi Goreng"
			hargaMakanan = 15000
		} else if pilihMakanan == "2" {
			namaMakanan = "Mie Goreng"
			hargaMakanan = 10000
		} else if pilihMakanan == "3" {
			namaMakanan = "Ayam Katsu"
			hargaMakanan = 20000
		} else if pilihMakanan == "4" {
			namaMakanan = "Sate"
			hargaMakanan = 15000
		} else if pilihMakanan == "5" {
			namaMakanan = "Ayam Balado"
			hargaMakanan = 17000
		} else {
			namaMakanan = "-"
			hargaMakanan = 0
			fmt.Println("angka yang di masukan tidak valid")
		}
	} else {
		namaMakanan = "-"
		hargaMakanan = 0
	}

	fmt.Print("Ingin memesan Minuman ? (iya/tidak) :")
	fmt.Scan(&mauMinuman)

	if mauMinuman == "iya" {
		fmt.Println("Menu Minuman :")
		fmt.Println("1. Lemon Tea | Rp : 10000")
		fmt.Println("2. Coklat    | Rp : 10000")
		fmt.Println("3. Kopi      | Rp : 5000")
		fmt.Println("4. Susu      | Rp : 5000")
		fmt.Println("5. Juice     | Rp : 10000")

		fmt.Print("Pilih Menu (1 - 4) :")
		fmt.Scan(&pilihMinuman)

		fmt.Print("jumlah Pesanan :")
		fmt.Scan(&jumlahMinuman)

		if pilihMinuman == "1" {
			namaMinuman = "Lemon Tea"
			hargaMinuman = 10000
		} else if pilihMinuman == "2" {
			namaMinuman = "Coklat"
			hargaMinuman = 10000
		} else if pilihMinuman == "3" {
			namaMinuman = "Kopi"
			hargaMinuman = 5000
		} else if pilihMinuman == "4" {
			namaMinuman = "Susu"
			hargaMinuman = 5000
		} else if pilihMinuman == "5" {
			namaMinuman = "Juice"
			hargaMinuman = 10000
		} else {
			namaMinuman = "-"
			hargaMinuman = 0
			fmt.Println("angka yang di masukan tidak valid")
		}
	} else {
		namaMinuman = "-"
		hargaMinuman = 0
	}

	fmt.Println("==========Pesanan Anda===============")
	fmt.Println("Makanan     :", namaMakanan)
	fmt.Println("Harga       : Rp.", hargaMakanan, "X", jumlahMakanan)
	fmt.Println("Minuman     :", namaMinuman)
	fmt.Println("Harga       : Rp.", hargaMinuman, "X", jumlahMinuman)
	totalHarga = (hargaMakanan * jumlahMakanan) + (hargaMinuman * jumlahMinuman)
	fmt.Println("=====================================")
	fmt.Println("Total Harga : Rp.", totalHarga)

	if totalHarga != 0 {
		cek := true
		for cek {
			fmt.Print("Masukan Nominal Pembayaran :")
			fmt.Scan(&pembayaran)

			if totalHarga <= pembayaran {
				kembalian = pembayaran - totalHarga
				fmt.Println("Kembalian : Rp.", kembalian)
				cek = false
			} else {
				fmt.Println("Uang yang anda masukan kurang")
			}
		}
	} else {
		fmt.Println("Belanjaan anda kosong")
	}

}
