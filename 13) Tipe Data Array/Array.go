package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "baskara"
	name[1] = "putra"
	name[2] = "hindia"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var value= [3]int{
		90,
		80,
		100,
	}

	fmt.Println(value)
	fmt.Println(len(value))
}