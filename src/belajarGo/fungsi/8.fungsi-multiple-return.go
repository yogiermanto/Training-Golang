package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("Luas Lingkaran\t\t: %.2f \n", area)
	fmt.Printf("Keliling Lingkaran\t: %.2f \n", circumference)
}

func calculate(d float64) (area float64, circumference float64) {
	//hitung luas
	area = math.Pi * math.Pow(d/2, 2)
	//hitung keliling
	circumference = math.Pi * d

	return

}
