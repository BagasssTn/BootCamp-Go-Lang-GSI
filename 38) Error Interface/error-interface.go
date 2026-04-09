package main

import (
	"errors"
	"fmt"
)

func main() {
	var angka1, angka2 int

	fmt.Print("Masukkan Angka Pertama: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan Angka Kedua: ")
	fmt.Scanln(&angka2)

	hasil, err := pembagian(angka1, angka2)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("tidak boleh dibagi 0")
	}
	result := nilai / pembagi
	return result, nil
}