package main

import (
	"fmt"
	"math/rand"
)

const SECONDS_PER_DAY = 60 * 60 * 24
const DISTANCE = 62100000

func main() {

	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")

	for i := 0; i < 10; i++ {
		var company string
		var trip string
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}
		speed := 16 + rand.Intn(15)
		days := DISTANCE / speed / SECONDS_PER_DAY

		price := 20 + speed

		if rand.Intn(2) == 0 {
			trip = "One-way"
		} else {
			trip = "Round-trip"
			price = price * 2
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", company, days, trip, price)
	}
}
