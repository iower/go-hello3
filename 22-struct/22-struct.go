package main

import (
	"encoding/json"
	"fmt"
)

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
		lat  float64 `json:"lat"`
		long float64 `json:"long"`
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

	// format field name
	fmt.Printf("%v\n", spirit2)
	fmt.Printf("%+ v\n", spirit2)

	// copy structs: by value
	spirit3 := spirit2
	spirit3.lat--
	fmt.Println(spirit2, spirit3)

	// struct slices
	locations := []location{
		{lat: 1, long: 2},
		{lat: 3, long: 4},
		{lat: 5, long: 6},
	}
	fmt.Println(locations)

	// struct -> json
	type loc2 struct {
		Lat, Long float64 // uppercase
	}
	loc2item := loc2{1, 2}
	bytes, err := json.Marshal(loc2item)
	if err == nil {
		fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	// tags
	type loc3 struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"lng"`
	}
	loc3item := loc3{3, 4}
	bytes2, err2 := json.Marshal(loc3item)
	if err2 == nil {
		fmt.Println(bytes2)
		fmt.Println(string(bytes2))
	}
}
