package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	fmt.Println(math.MaxInt64)
	var distance int64 = 41.3e12
	fmt.Println(distance)

	const lightSpeed = 299792
	const secondsPerDay = 86400

	days := distance / lightSpeed / secondsPerDay
	fmt.Println(days, "days")

	// big
	bigLightSpeed := big.NewInt(299792)
	bigSecondsPerDay := big.NewInt(86400)
	fmt.Println(bigLightSpeed, bigSecondsPerDay)

	bigDistance := new(big.Int)
	bigDistance.SetString("24000000000000000000", 10)

	bigSeconds := new(big.Int)
	bigSeconds.Div(bigDistance, bigLightSpeed)

	bigDays := new(big.Int)
	bigDays.Div(bigSeconds, bigSecondsPerDay)
	fmt.Println(bigDays, "days")

	// untyped int/float
	const untypedInt = 24000000000000000000
	fmt.Println("untypedInt", 24000000000000000000/299792/86400)
	const untypedFloat = untypedInt / 2.0
	fmt.Println("constDistance", untypedFloat)

	//
	kms := new(big.Int)
	kms.SetString("23600000000000000", 10)
	speed := new(big.Int)
	speed.SetInt64(299792)
	lightSeconds := new(big.Int)
	lightSeconds.Div(kms, speed)
	fmt.Println("lightSeconds", lightSeconds)

	secondsInYear := new(big.Int)
	secondsInYear.SetInt64(60 * 60 * 24 * 365)
	lightYears := new(big.Int)
	lightYears.Div(lightSeconds, secondsInYear)
	fmt.Println("lightYears", lightYears)
}
