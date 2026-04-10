package main

import (
	"fmt"
	"package-init/database"
)

func main()  {
	result := database.GetDatabase()
	fmt.Println(result)
}