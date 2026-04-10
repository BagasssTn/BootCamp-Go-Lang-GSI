package main

import (
	"fmt"
	"os" //ambil data pada sistem operasi/hardware
)

func main() {
	args := os.Args
	fmt.Println("Argumen : ")
	fmt.Println(args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}
}