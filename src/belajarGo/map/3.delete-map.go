package main

import "fmt"

func main() {
	var mahasiswa = map[string]string{
		"nama":    "yogi",
		"umur":    "22 tahun",
		"jurusan": "teknik informatika",
	}

	fmt.Println(len(mahasiswa))
	fmt.Println(mahasiswa)

	delete(mahasiswa, "umur")

	fmt.Println(len(mahasiswa))
	fmt.Println(mahasiswa)

}
