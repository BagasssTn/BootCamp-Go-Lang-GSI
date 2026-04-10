package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put your database host")
	var user *string = flag.String("user", "root", "put your database user")
	var pass *string = flag.String("pass", "localhost", "put your database pass")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Pass : ", *pass)
}