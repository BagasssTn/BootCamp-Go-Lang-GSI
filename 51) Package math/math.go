package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.609928))
	fmt.Println(math.Round(1.5))
	fmt.Println(math.Floor(1.72343)) //buat paksa bawah
	fmt.Println(math.Ceil(1.34443)) //bulat paksa atas

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}