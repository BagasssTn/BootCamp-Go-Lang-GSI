package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("b.*a") //validasi ada huruf b & a

	fmt.Println(regex.MatchString("baskara"))
	fmt.Println(regex.MatchString("baskoro"))
	fmt.Println(regex.MatchString("BISKIRI"))

	search := regex.FindAllString("eko eka edo eto eyo eki", -1)
	fmt.Println(search)
}