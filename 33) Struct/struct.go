package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var hindia Customer
	hindia.Name = "Baskara"
	hindia.Address = "Indonesia"
	hindia.Age = 30

	fmt.Println(hindia)

	mei := Customer {"Mei", "Indo", 30}
	fmt.Println(mei)
}