package aplikasi

import (
	"fmt"
	regis "mdb/model"
)

// Register is bla bla
func Register() {
	if regis.UmurSiswa < 40 && regis.BahasaSiswa == "Inggris" {
		fmt.Printf("Selamat %s Lulus", regis.NamaSiswa)
	} else {
		fmt.Printf("Maaf %s Tidak Lulus", regis.NamaSiswa)
	}
}
