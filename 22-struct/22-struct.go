package main

import "fmt"

func main() {
	// declare struct
	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -1.2345
	fmt.Println(curiosity)
	curiosity.long = 2.3456
	fmt.Println(curiosity)
	fmt.Println(curiosity.lat, curiosity.long)

	// reuse structs with types
	type location struct {
		lat  float64
		long float64
	}

	var spirit location
	spirit.lat = 1.2
	spirit.long = 1.22
	fmt.Println(spirit)

	var opportunity location
	opportunity.lat = 2.2
	opportunity.long = 2.22
	fmt.Println(opportunity)

	// init struct with composite literals
	type loc struct {
		lat, long float64
	}
	opportunity2 := loc{
		lat:  1234,
		long: 2345,
	}
	fmt.Println(opportunity2)

	// shorted declaraion
	spirit2 := loc{1.0, 2.0}
	fmt.Println(spirit2)
}
