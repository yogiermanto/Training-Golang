package main

import "fmt"

func main() {
	var (
		ijazah = "S1"
		bahasa = "Indonesia"
		umur   = 45
	)

	if ijazah == "S2" {
		fmt.Println("Ijazah anda Luar Biasa!")
	} else if ijazah == "S1" {
		fmt.Println("Ijazah anda Cukup!")
	} else {
		fmt.Println("Ijazah anda Kurang!")
	}

	if bahasa == "Inggris" {
		fmt.Println("Anda bisa bahasa Inggris")
	} else {
		fmt.Println("Anda harus bisa bahasa Inggris")
	}

	if umur > 40 {
		fmt.Println("Umur tidak bisa lebih dari 40")
	}

	for i := 1; i <= 20; i++ {
		fmt.Println("Angka", i)
	}

	// perulangan dengan argumen
	var i = 20
	for i <= 20 {
		fmt.Println("Angka", i)
		i++
	}

	// var ijazah string
	for {
		fmt.Println("Angka", i)
		i++
		if i == 20 {
			break
		}
	}
}
