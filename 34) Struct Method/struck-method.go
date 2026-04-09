package main

import "fmt"

type Customer struct {
	Nama, Addres string
	Age          int
}

func (a Customer) sayHii () {
	fmt.Println("Hiii from", a.Nama)
}

func (customer Customer) sayHey(name string) {
	fmt.Println("Hallo", name, "My Name is", customer.Nama)
}

func main() {
	var hindia Customer
	hindia.Nama = "Hindia"
	hindia.Addres = "Indonesia"
	hindia.Age = 30

	hindia.sayHey("Mei")
	hindia.sayHii()
}