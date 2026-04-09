package main

import "fmt"

func getFullName() (firstname string, midlename string, lastname string) {
	firstname = "Andrie"
	midlename = "Yunus"
	lastname = "Kurniawan"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(b)
}