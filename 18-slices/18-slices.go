package main

import "fmt"

func main() {
	strings := [...]string{"a", "b", "c", "d", "e", "f"}

	slice1 := strings[0:3]
	fmt.Println(slice1)
	fmt.Println(slice1[0], slice1[1], slice1[2])

	slice2 := strings[1:2]
	fmt.Println(slice2)

	// slice slice
	slice11 := slice1[1:2]
	fmt.Println(slice11)

	// slice changes when array changes
	strings[1] = "bb"
	fmt.Println(slice2, slice11)

	// default indexes
	strings0 := strings[:3]
	strings00 := strings[3:]
	stirngs000 := strings[:]
	fmt.Println(strings0, strings00, stirngs000)

	// apply for strings
	neptune := "Neptune"
	tune := neptune[3:] // new string
	fmt.Println(tune)
	neptune = "Newstrings"
	fmt.Println(tune)

	question := "¿Cómo estás?"
	fmt.Println(question[:6]) // "¿Cóm" - 6 bytes, not runes
}
