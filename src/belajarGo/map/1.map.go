package main

import "fmt"

func main() {
	var nilai map[string]int
	nilai = map[string]int{}

	nilai["lulus"] = 80
	nilai["gagal"] = 50

	fmt.Println("Nilai untuk Lulus : ", nilai["lulus"])
	fmt.Println("Nilai untuk Gagal : ", nilai["gagal"])
}
