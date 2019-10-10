package model

import "fmt"

var (
	//NamaSiswa bla bla
	NamaSiswa string = "Admin1"
	//UmurSiswa bla bla
	UmurSiswa int = 45
	//BahasaSiswa bla bla
	BahasaSiswa string = "Inggris"
)

// Siswa is representation of siswa
func Siswa() {

	fmt.Println(NamaSiswa)
	fmt.Println(UmurSiswa)
	fmt.Println(BahasaSiswa)
}
