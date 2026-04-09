package main

import "fmt"

func fullName() (string, string) {
	return "Andrie", "Yunus"
}
func main() {
	firstname, lastname := fullName()

	fmt.Println(firstname, lastname)
}