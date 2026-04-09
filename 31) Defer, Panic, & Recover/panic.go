package main

import "fmt"

func end() {
	message := recover() //Biar aplikasi tetap jalan
	if message != nil {
		fmt.Println("Error dengan :", message)
	}
	fmt.Println("Running")
}

func running(error bool)  {
	defer end()
	if error {
		panic("APLIKASI ERROR!!")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	running(false)
}