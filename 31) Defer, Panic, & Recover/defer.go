package main

import "fmt"

func logging() {
	fmt.Println("Disiram air keras aparat")
}

func runApp(string) {
	defer logging()
	fmt.Println("Siapa Andrie Yunus?")
	result := "Kenapa?"
	fmt.Println(result)
}

func main() {
	runApp("")
}