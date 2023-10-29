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

func mirepoix(ingredients []string) []string {
	fmt.Println(len(ingredients), cap(ingredients))
	return append(ingredients, "a", "b", "c")
}

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

type item string
type character struct {
	leftHand *item
}

func (c *character) pickup(i *item) {
	c.leftHand = i
}

func (c *character) give(to *character) {
	to.leftHand = c.leftHand
	c.leftHand = nil
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

	soup2 := mirepoix(nil)
	fmt.Println(soup2)

	// maps are nil by default
	var soup3 map[string]int
	fmt.Println(soup3 == nil)

	val, ok := soup3["onion"]
	fmt.Println(val, ok)

	for ingredient, measurement := range soup3 {
		fmt.Println(ingredient, measurement)
	}

	soup3 = map[string]int{"a": 1, "b": 2}
	fmt.Println(soup3)
	soup3 = nil
	fmt.Println(soup3, soup3 == nil)
	// soup3["onion"] = 1 // panic: assignment to entry in nil map

	// intrefaces are nil by default
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	fmt.Printf("%#v\n", v)

	var j *int
	v = j
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	fmt.Printf("%#v\n", v)

	//
	n := newNumber(42)
	fmt.Println(n)

	e := number{}
	fmt.Printf("%T \n", e)
	fmt.Println(e)

	var c1 character
	fmt.Println(c1)
	var sword item = "sword"
	c1.pickup(&sword)
	fmt.Println(c1)

	var c2 character
	fmt.Println(c2)
	c1.give(&c2)
	fmt.Println(c1, c2)
}
