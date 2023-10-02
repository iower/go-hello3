package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)

	/*
		// errors
		&42
		&"indirection"
		&true
	*/

	address := &answer
	fmt.Println(*address)

	/*
		address++
	*/
	fmt.Println(*&*&answer)

	// type: *int
	fmt.Printf("address type is %T\n", address)

	canada := "Canada"
	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &canada
	fmt.Println(*home)
}
