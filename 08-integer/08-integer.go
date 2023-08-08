package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	year := 2023     // int
	var year2 = 2018 // int
	var year3 int = 2023
	var month uint = 2

	days := 365.2425

	fmt.Printf("Types: %T, %T, %T, %T, %T \n", year, year2, year3, month, days)

	// hex

	var red1, green1, blue1 uint8 = 0, 141, 213
	var red2, green2, blue2 uint8 = 0x00, 0x8d, 0xd5

	fmt.Printf("%x %x %x \n", red1, green1, blue1)
	fmt.Printf("%x %x %x \n", red2, green2, blue2)

	fmt.Printf("color: #%02x%02x%02x\n", red1, green1, blue1)

	// overflow
	fmt.Println()

	var num uint8 = 255
	fmt.Printf("%v %08b \n", num, num)
	num++
	fmt.Printf("%v %08b \n", num, num)

	var num2 int8 = 127
	fmt.Printf("%v %08b \n", num2, num2)
	num2++
	fmt.Printf("%v %08b \n", num2, num2)

	var num3 uint8 = 255
	fmt.Printf("%v %08b \n", num3, num3)
	num3++
	fmt.Printf("%v %08b \n", num3, num3)

	// math.Max...
	fmt.Println()
	fmt.Println(math.MaxInt8, math.MaxUint8)
	fmt.Println(math.MaxInt16, math.MaxUint16)
	fmt.Println(math.MaxInt32, math.MaxUint32)
	var m uint64 = math.MaxUint64
	fmt.Println(math.MaxInt64, m)
	fmt.Println(math.MaxInt)

	//

	var pig int = 0
	for pig < 25*100 {
		switch rand.Intn(3) {
		case 0:
			pig += 5
		case 1:
			pig += 10
		case 2:
			pig += 125
		}
		fmt.Printf("Balance: $%d.%02v \n", pig/100, pig%100)
	}
}
