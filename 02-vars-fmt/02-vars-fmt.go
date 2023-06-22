package main

import "fmt"

func main() {
	fmt.Print("Weight: ")
	fmt.Print(55 * 0.3783)
	fmt.Print(" kg, Age: ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years.")

	fmt.Println()

	fmt.Printf("Weight: %v, Age: %v\n", 55*0.3783, 41*365/687)
}
