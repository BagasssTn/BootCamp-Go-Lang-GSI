package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Hindia",
		"job":  "Singer",
	}

	person["title"] = "Goat"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["job"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Mencintai"
	book["author"] = "Betri"
	book["year"] = "2003"
	fmt.Println(book)

	delete(book, "year")

	fmt.Println(book)
}
