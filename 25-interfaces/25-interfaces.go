package main

import (
	"fmt"
	"strings"
	"time"
)

var t interface {
	talk() string
}

//

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

// -er rule

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type crater struct{}

type starship struct {
	laser
}

type rover struct{}

func (r rover) talk() string {
	return "whir whir"
}

type stardater interface {
	YearDay() int
	Hour() int
}

// func stardate(t time.Time) float64 {
func stardate(t stardater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + day + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

func main() {
	fmt.Println(t)

	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(2))

	// shout(crater{})

	s := starship{laser(4)}
	fmt.Println(s.talk())
	shout(s)

	r := rover{}
	fmt.Println(r.talk())
	shout(r)

	//
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	ss := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(ss))
}
