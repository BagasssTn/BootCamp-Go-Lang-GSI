package main

import "fmt"

func main() {
	var name = "MBG"

	if name == "Embege" {
		fmt.Println("Embege Jelek")
	} else if name == "MBG" {
		fmt.Println("Iyaa")
	} else {
		fmt.Println("Iya, Embege Jelek")
	}

	if length := len(name) ; length > 5 {
		fmt.Println("Terlalu mahal")
	} else {
		fmt.Println("Tetap Mahal")
	}
}