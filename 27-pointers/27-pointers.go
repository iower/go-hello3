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

	//

	var admin *string
	alice := "Alice"
	admin = &alice
	fmt.Println(*admin)

	bob := "Bob"
	admin = &bob
	fmt.Println(*admin)

	bob = "Bob2"
	fmt.Println(*admin)

	*admin = "Bob3"
	fmt.Println(*admin, bob)

	major := admin
	*major = "Bob4"
	fmt.Println(*admin, *major, admin == major)

	newMajor := "newMajor"
	major = &newMajor
	fmt.Println(*admin, *major, admin == major)

	//
	strClone := *major
	*major = "newMajor2"
	fmt.Println(strClone, *major)

	testEquality := "newMajor2"
	fmt.Println(newMajor, &newMajor, testEquality, &testEquality)
	fmt.Println(newMajor == testEquality)
	fmt.Println(&newMajor == &testEquality)

	// pointers and structs
	type person struct {
		name, superpower string
		age              int
	}

	timmy := &person{
		name: "Tim",
		age:  20,
	}
	fmt.Println(timmy)
	// shorter than
	// (*timmy).superpower = "flying"
	timmy.superpower = "flying"
	fmt.Println(timmy)
	fmt.Printf("%+v\n", timmy)
	fmt.Println(timmy.superpower, (*timmy).superpower)
}
