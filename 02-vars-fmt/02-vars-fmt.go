package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Weight: ")
	fmt.Print(55 * 0.3783)
	fmt.Print(" kg, Age: ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years.")

	fmt.Println()

	fmt.Printf("Weight: %v, Age: %v\n", 55*0.3783, 41*365/687)

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	const lightSpeed = 299792
	var distance = 56000000

	fmt.Println(distance/lightSpeed, "s")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "s")

	var (
		distance2 = 56000000
		speed2    = 100800
	)
	fmt.Println(distance2 / speed2)

	var distance3, speed3 = 56000000, 100800
	fmt.Println(distance3 / speed3)

	var weight = 149.0
	weight *= 0.3783

	var age = 41
	age++

	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
