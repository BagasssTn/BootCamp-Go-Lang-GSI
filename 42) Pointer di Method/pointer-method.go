package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {  //Man harus pake * 
	man.Name = "Mr " + man.Name
}

func main() {
	person := Man{"Hindia"}
	person.Married()

	fmt.Println(person.Name)
}