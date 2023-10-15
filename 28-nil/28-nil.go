package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	p.age++
}

func (p *person) birthdayImproved() {
	if p == nil {
		return
	}
	p.age++
}

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

	var nobody *person
	fmt.Println(nobody)
	//	nobody.birthday() // error: nil pointer dereference
	fmt.Println(nobody)
	nobody.birthdayImproved()
	fmt.Println(nobody)
	// fmt.Println(nobody.age) // error: nil pointer dereference

	// funcs are nil by default
	var fn func(a, b int) int
	fmt.Println(fn == nil)
	// fn(1, 2) // error: nil pointer dereference

}
