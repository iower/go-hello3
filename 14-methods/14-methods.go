package main

import "fmt"

type celsius float64
type fahrenheith float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelving(c celsius) kelvin {
	return kelvin(c + 273.15)
}

// methods

func (k kelvin) toCelsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheith) toCelsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var temperature celsius = 10.0
	fmt.Println(temperature)
	temperature += 10
	fmt.Println(temperature)

	// const degrees = 30.0
	// temperature := degrees
	// fmt.Println(temperature)

	// var warmUp float64 = 10.0
	// temperature += warmUp
	// fmt.Println(temperature)

	// convert to type
	var warmUp2 float64 = 10
	temperature += celsius(warmUp2)

	//
	var c celsius = 20
	var f fahrenheith = 20

	// if c == f { } // mismatched types
	// c += f // mismatched types
	fmt.Println(c, f)

	//
	var k kelvin = 294
	cc := kelvinToCelsius(k)
	fmt.Println(k, "° K is", cc, "° C")

	// use methods
	ccc := k.toCelsius()
	fmt.Println(k, "° K is", ccc, "° C")
}
