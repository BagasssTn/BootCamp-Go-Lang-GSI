package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Nama": name,
		}
	}
}
func main() {
	var person map[string]string = newMap("Baskara")

	if person == nil {
		fmt.Println("Data Khosong")
	} else {
		fmt.Println(person)
	}

}
