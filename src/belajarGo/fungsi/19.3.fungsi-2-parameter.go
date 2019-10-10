package main

import (
	"fmt"
	"strings"
)

func main() {
	hobiSaya("Yogi", "Gaming", "Ngoding")
}

func hobiSaya(nama string, hobi ...string) {
	var hobiSbgString = strings.Join(hobi, ", ")

	fmt.Printf("Halo, nama saya adalah: %s\n", nama)
	fmt.Printf("Hobi saya adalah: %s\n", hobiSbgString)
}
