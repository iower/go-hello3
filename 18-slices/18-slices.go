package main

import (
	"fmt"
	"strings"
)

func main() {
	strings_ := [...]string{"a", "b", "c", "d", "e", "f"}

	slice1 := strings_[0:3]
	fmt.Println(slice1)
	fmt.Println(slice1[0], slice1[1], slice1[2])

	slice2 := strings_[1:2]
	fmt.Println(slice2)

	// slice slice
	slice11 := slice1[1:2]
	fmt.Println(slice11)

	// slice changes when array changes
	strings_[1] = "bb"
	fmt.Println(slice2, slice11)

	// default indexes
	strings0 := strings_[:3]
	strings00 := strings_[3:]
	stirngs000 := strings_[:]
	fmt.Println(strings0, strings00, stirngs000)

	// apply for strings
	neptune := "Neptune"
	tune := neptune[3:] // new string
	fmt.Println(tune)
	neptune = "Newstrings"
	fmt.Println(tune)

	question := "¿Cómo estás?"
	fmt.Println(question[:6]) // "¿Cóm" - 6 bytes, not runes

	// slice from array
	arr2 := [...]string{"aaa", "bbb", "ccc"}
	sliceFromArr := arr2[:]
	fmt.Println(sliceFromArr)

	// composite literals
	slice3 := []string{"aaa", "bbb", "ccc"}
	fmt.Println(slice3)

	fmt.Printf("%T, %T, %T\n", arr2, sliceFromArr, slice3)

	//
	slice4 := []string{" Aa ", " Bb   ", "\nCc\n"}
	trim(slice4)
	fmt.Println(strings.Join(slice4, ""))
}

func trim(strs []string) { // any slice length
	for i := range strs {
		strs[i] = strings.TrimSpace(strs[i])
	}
}
