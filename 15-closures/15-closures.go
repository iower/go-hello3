package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func measureTemperature(samples int, sensor func() kelvin) {
	// measureTemperature(samples int, s sensor)
	// var sensor func() kelvin
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func main() {
	// first-class func
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	var sensor2 func() kelvin
	fmt.Println(sensor, sensor2)

	//
	measureTemperature(3, fakeSensor)

	f()
}

/*

func drawTable(rows int, getRow func(row int) (string, string))
->
type getRowFn func(row int) (string, string)
func drawTable(rows int, getRow getRowFn)
*/

var f = func() {
	fmt.Println("Anonymous function")
}
