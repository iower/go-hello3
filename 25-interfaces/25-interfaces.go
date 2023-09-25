package main

import "fmt"

func main() {
	var t interface {
		talk() string
	}

	fmt.Println(t)
}
