package main

import "fmt"

func main() {
	var nama, kelas, fakultas, jurusan string = "Yogi", "4IA16", "Teknologi Industri", "Teknik Informatika"
	var npm uint32 = 57415251
	var _ = "22 tahun"

	fmt.Println("Mahasiswa")
	fmt.Printf("Nama : %s\n", nama)
	fmt.Printf("NPM : %d\n", npm)
	fmt.Printf("Kelas : %s\n", kelas)
	fmt.Printf("Fakultas : %s\n", fakultas)
	fmt.Printf("Jurusan : %s\n", jurusan)

}
