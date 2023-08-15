package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	countdown := "Launch in T minus " + "10 seconds"
	fmt.Println(countdown)

	// invalid
	// countdown2 := "Launch in T minus " + 10 + " seconds"

	// int to float
	age := 41
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.")

	// float to int
	fmt.Println(int(earthDays))

	// overflow
	red := 256
	fmt.Println(uint8(red)) // 0

	var floatValue float64 = 32768
	var intValue = int16(floatValue) // max 32767
	fmt.Println(intValue)            // -32768

	// prevent overflow

	var floatValue2 float64 = 32767
	if floatValue2 < math.MinInt16 || floatValue2 > math.MaxInt16 {
		panic("Out of range")
	} else {
		var intValue2 = int16(floatValue2) // max 32767
		fmt.Println(intValue2)
	}

	// check 8bit uint
	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("Converted:", v8)
	}

	// rune to string
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))

	// integer to ASCII
	countdown2 := 10
	str2 := "Launch in T minus " + strconv.Itoa(countdown2) + " seconds"
	fmt.Println(str2)
}
