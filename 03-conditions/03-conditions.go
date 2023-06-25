package main

import (
	"fmt"
	"strings"
)

func main() {
	var command = "go outside"
	var isExit = strings.Contains(command, "outside")
	fmt.Println(isExit)

	var age = 41
	var minor = age < 18
	fmt.Printf("Is minor %v\n", minor)

	fmt.Println("Apple" > "banana")
	fmt.Println("apple" > "Banana")

	var year = 2000
	var leap = year%400 == 0 || (year&4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("leap")
	} else {
		fmt.Println("not leap")
	}
}
