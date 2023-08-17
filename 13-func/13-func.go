package main

import (
	"fmt"
	"strconv"
	"time"
)

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheith(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheith(kelvinToCelsius(k))
}

func main() {
	future := time.Unix(12622780800, 0)
	fmt.Println(future)

	// multiple return
	countdown, err := strconv.Atoi("10")
	fmt.Println(countdown, err)

	//
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "° K is ", celsius, "° C\n")

	kelvin2 := 233.0
	celsius2 := kelvinToCelsius(kelvin2)
	fmt.Print(kelvin2, "° K is ", celsius2, "° C\n")
}
