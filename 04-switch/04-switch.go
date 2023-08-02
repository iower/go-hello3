package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	weekday := now.Weekday()

	switch weekday {
	case time.Monday:
		fmt.Println("Today is Monday")
	case time.Tuesday:
		fmt.Println("Today is Tuesday")
	case time.Wednesday:
		fmt.Println("Today is Wednesday")
	case time.Thursday:
		fmt.Println("Today is Thursday")
	case time.Friday:
		fmt.Println("Today is Friday")
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	}

	// few cases, default case
	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("(day of rest)")
	default:
		fmt.Println("(workday)")
	}

	// initializator
	switch num := 6; num%2 == 0 {
	case true:
		fmt.Println(num, "- even")
	case false:
		fmt.Println(num, "- odd")
	}

	// break
	str := "a b c\td\nefg hi"
	for _, letter := range str {
		switch letter {
		case ' ', '\t', '\n':
			break
		default:
			fmt.Printf("%c\n", letter)
		}
	}

	// switch true
	switch {
	case now.Hour() < 12:
		fmt.Println("AM")
	default:
		fmt.Println("PM")
	}

	// fallthrough
	nextstop := "B"
	fmt.Println("Stops adead of us:")

	switch nextstop {
	case "A":
		fmt.Println("A")
		fallthrough
	case "B":
		fmt.Println("B")
		fallthrough
	case "C":
		fmt.Println("C")
		fallthrough
	case "D":
		fmt.Println("D")
		fallthrough
	case "E":
		fmt.Println("E")
	}

	// switch type
	var data interface{}
	data = 123456789.123456789

	switch mytype := data.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case int:
		fmt.Println("int")
	default:
		fmt.Printf("%T", mytype)
	}
}
