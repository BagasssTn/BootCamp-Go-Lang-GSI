package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Daniel Baskara", "Daniel"))
	fmt.Println(strings.Contains("Daniel Baskara", "Putra"))

	fmt.Println(strings.Split("Daniel Baskara Putra", " "))

	fmt.Println(strings.ToLower("Daniel Baskara Putra"))
	fmt.Println(strings.ToUpper("DaNiel BaSkara Putra"))
	fmt.Println(strings.ToTitle("daniel baskara putra"))

	fmt.Println(strings.Trim("      Daniel Baskara     ", " "))
}