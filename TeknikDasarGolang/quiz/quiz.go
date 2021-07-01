package main

import (
	"fmt"
	"quiz/multiplication"
)

func main() {
	fmt.Println("Belajar Perkalian Menggunakan Golang")

	result := multiplication.Kali(5, 5)

	fmt.Println(result)
}
