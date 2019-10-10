package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = fruits[0:2]

	fmt.Println(fruits)
	fmt.Println(bFruits, "Cap :", cap(bFruits), "Len :", len(bFruits))

	fmt.Printf("\nDi append jika slice len < cap\n\n")

	var cFruits = append(bFruits, "pepaya")

	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
}
