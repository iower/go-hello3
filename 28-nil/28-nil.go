package main

import "fmt"

func main() {
	var i int
	fmt.Println(i)
	var s string
	fmt.Println(s)

	var p *int // pointers are nil by default
	fmt.Println(p)

	var ii interface{} // interfaces are nil by default
	fmt.Println(ii)

	//
	fmt.Println("---")
	var nowhere *int
	fmt.Println(nowhere)
	// fmt.Println(*nowhere) // error: nil pointer dereference
	if nowhere != nil {
		fmt.Println(*nowhere)
	}
}
