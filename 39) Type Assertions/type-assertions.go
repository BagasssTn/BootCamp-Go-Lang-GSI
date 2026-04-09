package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Interger")
	default:
		fmt.Println("Unokwon Type")
	} 
}