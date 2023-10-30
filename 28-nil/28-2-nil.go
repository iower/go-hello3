package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v pickups %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v cannot give\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v cannot take\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	fmt.Printf("%v gives %v %v\n", c.name, to.name, to.leftHand.name)
	c.leftHand = nil
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v noting carries", c.name)
	}
	return fmt.Sprintf("%v carries %v", c.name, c.leftHand.name)
}

func main() {
	alice := &character{name: "Alice"}
	someItem := &item{name: "SomeItem"}
	alice.pickup(someItem)

	bob := &character{name: "Bob"}
	alice.give(bob)

	fmt.Println(alice)
	fmt.Println(bob)
}
