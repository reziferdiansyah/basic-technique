package main

import "fmt"

func main() {
	title := "GOLANG THE BEST LANGUAGE"

	for index, letter := range title {
		fmt.Println("index:", index, "letter :", string(letter))
	}

	//jika tidak menggunakan index

	// title := "GOLANG THE BEST LANGUAGE"
	// for _, letter := range title {
	// 	fmt.Println("letter :", string(letter))
	// }
}
