package main

import "fmt"

func main() {
	var buah = []string{"anggur", "apel", "rambutan"}
	var aBuah = buah[0:2]
	var bBuah = buah[0:2:2]

	fmt.Println("Slice : ", buah, "Len : ", len(buah), "Cap : ", cap(buah))

	fmt.Println("Slice : ", aBuah, "Len : ", len(aBuah), "Cap : ", cap(aBuah))

	fmt.Println("Slice : ", bBuah, "Len : ", len(bBuah), "Cap : ", cap(bBuah))
}
