package main

import (
	"fmt"
	"sort"
)

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

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return s[i] < s[j]
		}
	}
	sort.Slice(s, less)
}

func sortStringsByLen(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return len(s[i]) < len(s[j])
		}
	}
	sort.Slice(s, less)
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

	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)
	sortStringsByLen(food, nil)
	fmt.Println(food)

	// slices are nil by default
	var soup []string
	fmt.Println(soup == nil)

	for _, i := range soup {
		fmt.Println(i)
	}

	fmt.Println(len(soup)) // 0
	soup = append(soup, "a", "b", "c")
	fmt.Println(soup)
}
