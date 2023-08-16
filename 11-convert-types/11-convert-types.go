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

	// ASCII to integer
	countdown3, err := strconv.Atoi("10")
	if err != nil {
		panic(err)
	}
	fmt.Println(countdown3)

	// Sprintf
	str3 := fmt.Sprintf("Launch in T minus %v seconds", countdown2)
	fmt.Println(str3)

	// static typing
	var staticInt = 10
	// staticInt = 10.1 // error
	// staticInt = fmt.Sprintf("%v seconds", staticInt) // error
	fmt.Println(staticInt)

	// boolean to string
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)

	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)

	// string(false) // error

	// string to boolean
	launch2 := yesNo == "yes"
	fmt.Println("Ready for launch:", launch2)

	// boolean to int
	boolVal := false
	var intVal int
	if boolVal {
		intVal = 1
	} else {
		intVal = 0
	}
	fmt.Println(intVal)

	//
	str := "true"
	var boolFromStr bool

	switch str {
	case "true", "yes", "1":
		boolFromStr = true
	case "false", "no", "0":
		boolFromStr = false
	default:
		panic("Unknown str")
	}
	fmt.Println(boolFromStr)
}
