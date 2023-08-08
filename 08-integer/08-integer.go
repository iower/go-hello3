package main

import "fmt"

func main() {
	year := 2023     // int
	var year2 = 2018 // int
	var year3 int = 2023
	var month uint = 2

	days := 365.2425

	fmt.Printf("Types: %T, %T, %T, %T, %T \n", year, year2, year3, month, days)

	// hex

	var red1, green1, blue1 uint8 = 0, 141, 213
	var red2, green2, blue2 uint8 = 0x00, 0x8d, 0xd5

	fmt.Printf("%x %x %x \n", red1, green1, blue1)
	fmt.Printf("%x %x %x \n", red2, green2, blue2)

	fmt.Printf("color: #%02x%02x%02x\n", red1, green1, blue1)

	// overflow
	// ...
}
