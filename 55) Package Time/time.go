package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2026, 12, 13, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2026-12-13"
	parse, _ := time.Parse(layout, "1990-03-20")
	fmt.Print(parse)
}