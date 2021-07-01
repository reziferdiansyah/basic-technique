// Package Executable

package main

import (
	"fmt"
	"kedua/calculation"
)

func main() {
	fmt.Println("Belajar Pertambahan Menggunakan Golang")

	result := calculation.Add(5, 5)

	fmt.Println(result)
}
