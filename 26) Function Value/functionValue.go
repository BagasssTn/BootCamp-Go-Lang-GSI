package main

import "fmt"

func awoo(name string) string {
	return "Hawoo " + name + "!"
}

func main() {

	sayAwoo := awoo

	result := sayAwoo("Andrie")

	fmt.Println(result)
}