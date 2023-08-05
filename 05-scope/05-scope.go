package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

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

	//
	for i := 0; i < 10; i++ {
		generateDate()
	}
}

func generateDate() {
	year := rand.Intn(2023) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		isLeap := (year%4 == 0 && year%100 != 0) || year%400 == 0
		if isLeap {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 1:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}
