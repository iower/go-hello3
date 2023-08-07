package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	days := 365.2425
	var days2 = 365.2425
	var days3 float64 = 365.2425

	fmt.Println(days, days2, days3)

	// precision
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64, pi32)

	// initial value
	var price float64
	fmt.Println(price)

	// format
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
	fmt.Printf("%06.2f\n", third)

	// rounding error
	fmt.Println(third + third + third)

	bank := 0.1
	bank += 0.2
	fmt.Println(bank)
	fmt.Println(bank == 0.3)
	fmt.Println(math.Abs(bank - 0.3))

	// multiplication before division
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "° F\n")
	fmt.Print((9.0/5.0*celsius)+32, "° F\n")

	fmt.Print((celsius*9.0/5.0)+32, "° F\n")

	//
	for piggy := 0.0; piggy < 20; {
		switch rand.Intn(3) {
		case 0:
			piggy += 0.05
		case 1:
			piggy += 0.10
		case 2:
			piggy += 0.25
		}
		fmt.Printf("%5.2f\n", piggy)
	}
}
