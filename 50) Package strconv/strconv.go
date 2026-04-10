package main

import (
	"fmt"
	"strconv" //string conversion (string to int or int to string)
)

func main() {
	boolean, err := strconv.ParseBool("Benar")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 8)
	fmt.Println(value)

	valueInt, _ :=  strconv.Atoi("2000000")
	fmt.Println(valueInt)

	valuye := 2000
	str := strconv.Itoa(valuye)
	fmt.Println("ini string :", str)	
}