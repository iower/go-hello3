package main

import "fmt"

type converter func(a float64) float64

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32.0) * 5.0 / 9.0
}

func drawTable(title1, title2 string, converterFunc converter) {
	fmt.Println("=======================")
	fmt.Printf("| %-8s | %-8s |\n", title1, title2)
	fmt.Println("=======================")
	for t := 40; t <= 100; t += 5 {
		f := float64(t)
		fmt.Printf("| %-8.2f | %-8.2f |\n", f, converterFunc(f))
	}
	fmt.Println("=======================")
}

func main() {
	drawTable("°C", "°F", celsiusToFahrenheit)
	drawTable("°F", "°C", fahrenheitToCelsius)
}
