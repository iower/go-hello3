package main

import "fmt"

func main() {
	// append to slice
	strings := []string{"a", "b", "c"}
	fmt.Println(strings, len(strings), cap(strings))

	strings = append(strings, "d")
	fmt.Println(strings, len(strings), cap(strings))

	strings = append(strings, "e", "f")
	fmt.Println(strings, len(strings), cap(strings))

	more := []string{"g", "h"}
	strings = append(strings, more...)
	fmt.Println(strings, len(strings), cap(strings))
}
