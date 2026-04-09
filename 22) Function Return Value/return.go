package main

import "fmt"

func airKeras(name string) string {
	if name == "" {
		return "Hawoo Bruuh"
	} else {
		return "Hawoo " + name
	}
}
func main() {
	result := airKeras("Andrie")
	fmt.Println(result)

	fmt.Println(airKeras(""))
}