package main

import "fmt"

func main() {
	strings := [...]string{"a", "b", "c", "d", "e", "f"}

	slice1 := strings[0:3]
	fmt.Println(slice1)
	fmt.Println(slice1[0], slice1[1], slice1[2])

	slice2 := strings[1:2]
	fmt.Println(slice2)

}
