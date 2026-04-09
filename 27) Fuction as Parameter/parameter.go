package main

import "fmt"

type Filter func(string) string

func sayHelloFilter (name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Bahlil" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("Andrie", spamFilter)
	sayHelloFilter("Bahlil", spamFilter)
}