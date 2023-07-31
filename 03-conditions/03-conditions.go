package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var command = "go outside"
	var isExit = strings.Contains(command, "outside")
	fmt.Println(isExit)

	var age = 41
	var minor = age < 18
	fmt.Printf("Is minor %v\n", minor)

	fmt.Println("Apple" > "banana")
	fmt.Println("apple" > "Banana")

	var year = 2000
	var leap = year%400 == 0 || (year&4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("leap")
	} else {
		fmt.Println("not leap")
	}

	// switch
	var switchCommand = 3
	switch switchCommand {
	case 1:
		fmt.Println("Command 1")
	case 2:
		fmt.Println("Command 2")
	default:
		fmt.Println("Other")
		fallthrough
	case 1000:
		fmt.Println("Fallthrough")
	}

	// for
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("The end!")

	// endless for
	var degrees = 0

	for {
		fmt.Println(degrees)

		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}

	//
	for counter := 10; counter >= 0; counter-- {
		if rand.Intn(100) == 0 {
			fmt.Println("Break o_o")
			break
		} else {
			fmt.Println(counter)
			time.Sleep(time.Second)
		}
	}
}
