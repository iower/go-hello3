package main

import "fmt"

func main() {
	// append to slice
	strings := []string{"a", "b", "c"}
	strings = append(strings, "d")

	fmt.Println(strings)
	strings = append(strings, "e", "f")
	fmt.Println(strings)

	more := []string{"g", "h"}
	strings = append(strings, more...)
}
