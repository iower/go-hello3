package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		{
			sign = -1
		}
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// constructor: newType
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal())

	curiosity := location{lat.decimal(), long.decimal()}
	fmt.Println(curiosity)

	curiosity2 := newLocation(coordinate{5, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity2)

	// classes
	var mars = world{radius: 3389.5}
	fmt.Println(mars)
	dist := mars.distance(curiosity, curiosity2)
	fmt.Printf("%.2f km \n", dist)

	//
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	dist2 := mars.distance(spirit, opportunity)
	fmt.Println(dist2)

	//
	earth := world{radius: 6371}
	london := newLocation(coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	paris := newLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	fmt.Println(earth.distance(london, paris))
}
