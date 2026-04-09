package main

import "fmt"

func main() {
	name := "Bagas"

	switch name {
	case "Bagas":
		fmt.Println("Bagas Ganteng")
	
	case "Tri":
		fmt.Println("Tricahya")

	default:
		fmt.Println("Hallo Gantenk")
	}

		// switch length := len(name); length > 5 {
		// case true:
		// 	fmt.Println("Nama Terlalu Panjang")
		// case false:
		// 	fmt.Println("Nama Sudah Benar")
		// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Cocok")
	default:
		fmt.Println("Nama Benar")
	}
}