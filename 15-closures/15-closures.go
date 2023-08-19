package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func main() {
	// first-class func
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	var sensor2 func() kelvin
	fmt.Println(sensor, sensor2)
}
