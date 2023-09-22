package main

import "fmt"

type report0 struct {
	sol       int
	high, low float64
	lat, long float64
}

// composition

type temperature struct {
	high, low celsius
}

type celsius float64

type location struct {
	lat, long float64
}

type report struct {
	sol         int
	temperature temperature
	location    location
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (r report) averageT() celsius {
	return r.temperature.average()
}

// method forwarding
type report2 struct {
	sol
	temperature
	location
}

type sol int

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v*C\n", report.temperature.high)
	fmt.Println("Average t: ", report.temperature.average())
	fmt.Println("Average t: ", t.average())
	fmt.Println("Average t: ", report.averageT())

	// method forwarding
	report2 := report2{sol: 15, temperature: t, location: bradbury}
	fmt.Println("Forwarded method", report2.average(), report2.days(10))
	fmt.Println("Forwarded field", report2.high)
}
