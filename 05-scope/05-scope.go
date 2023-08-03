package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	for count < 10 { // scope start
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	} // scope end
	// fmt.Println(num) // out of scope

	//

	var count2 = 0
	for count2 = 10; count2 > 0; count2-- {
		fmt.Println(count2)
	}
	fmt.Println(count2) // in scope

	//

	for count3 := 10; count3 > 0; count3-- {
		fmt.Println(count3)
	}
	// fmt.Println(count3) // out of scope

	//

	if num := rand.Intn(3); num == 0 {
		fmt.Println("num == 0")
	} else if num == 1 {
		fmt.Println("num == 1")
	} else {
		fmt.Println("num > 1")
	}
	// fmt.Println(num) // out of scope

	//

	switch num := rand.Intn(3); num {
	case 0:
		fmt.Println("num == 0")
	case 1:
		fmt.Println("num == 1")
	default:
		fmt.Println("num > 1")
	}
	// fmt.Println(num) // out of scope
}
