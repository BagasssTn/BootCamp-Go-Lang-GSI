package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountry(adress *Address) { //Kalau pake * negaranya berubah
	adress.Country = "Konoha"
}

func main() {
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Lamongan"
	address2 = &Address{"Malang", "Jawa Timur", "Indonesia"} //Pass by reference tanpa operator
	//*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} //Dengan Operator
	fmt.Println(address1) //Tanpa tidak berubah, dengan berubah
	fmt.Println(address2)

	var address4 *Address = new(Address) //Membuat pointer baru
	fmt.Println(address4)

	var alamat = Address {
	City: "Malang",
	Province: "Jawa Timur",
	Country: "",
	}

	ChangeCountry(&alamat) //pake & berubah
	fmt.Println(alamat)
}