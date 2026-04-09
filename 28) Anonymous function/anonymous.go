package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are dictactor", name)
	} else {
		fmt.Println("You are normal person", name)
	}
}

// func blacklist1(name string) bool {
// 	return name == "Wowo"
// }

// func blacklistRoot(name string) bool {
// 	return name == "Andrie"
	
// }

func main() {
	BlackList := func (name string) bool {
		return name == "wowo"
	}

	registerUser("wowo", BlackList)
	registerUser("andrie", BlackList)

	registerUser("bahlil", func (name string) bool {
		return name == "bahlil"
	})
	registerUser("teddy", func (name string) bool {
		return name == "teddy"
	})
}