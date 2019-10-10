package main

import "fmt"

func main() {
	var mahasiswa = map[string]string{
		"nama":    "yogi",
		"umur":    "22 tahun",
		"jurusan": "teknik informatika",
	}

	for key, val := range mahasiswa {
		fmt.Println(key, " :", val)
	}

}
