package model

import "fmt"

type murid struct {
	nama     string
	jurusan  string
	angkatan int
}

//StructMurid adalah representasi dari struct
func StructMurid() {
	var s1 = murid{}
	s1.nama = "Yogi"
	s1.jurusan = "Teknik Informatika"
	s1.angkatan = 2015

	var s2 = murid{"Ermanto", "Teknik Mesin", 2016}

	var s3 = murid{nama: "Farrel"}

	fmt.Println("murid 1 :", s1.nama)
	fmt.Println("murid 2 :", s2.nama)
	fmt.Println("murid 3 :", s3.nama)
}
