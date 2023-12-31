package main

import (
	"fmt"
	"strings"
)

type person struct {
	name, superpower string
	age              int
}

// pointers as params
func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

// pointers and interfaces
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.x++
}
func (t *turtle) down() {
	t.x--
}
func (t *turtle) left() {
	t.y--
}
func (t *turtle) right() {
	t.y++
}
func (t turtle) where() {
	fmt.Printf("turtle (%v, %v)\n", t.x, t.y)
}

func main() {
	answer := 42
	fmt.Println(&answer)

	/*
		// errors
		&42
		&"indirection"
		&true
	*/

	address := &answer
	fmt.Println(*address)

	/*
		address++
	*/
	fmt.Println(*&*&answer)

	// type: *int
	fmt.Printf("address type is %T\n", address)

	canada := "Canada"
	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &canada
	fmt.Println(*home)

	//

	var admin *string
	alice := "Alice"
	admin = &alice
	fmt.Println(*admin)

	bob := "Bob"
	admin = &bob
	fmt.Println(*admin)

	bob = "Bob2"
	fmt.Println(*admin)

	*admin = "Bob3"
	fmt.Println(*admin, bob)

	major := admin
	*major = "Bob4"
	fmt.Println(*admin, *major, admin == major)

	newMajor := "newMajor"
	major = &newMajor
	fmt.Println(*admin, *major, admin == major)

	//
	strClone := *major
	*major = "newMajor2"
	fmt.Println(strClone, *major)

	testEquality := "newMajor2"
	fmt.Println(newMajor, &newMajor, testEquality, &testEquality)
	fmt.Println(newMajor == testEquality)
	fmt.Println(&newMajor == &testEquality)

	// pointers and structs

	timmy := &person{
		name: "Tim",
		age:  20,
	}
	fmt.Println(timmy)
	// shorter than
	// (*timmy).superpower = "flying"
	timmy.superpower = "flying"
	fmt.Println(timmy)
	fmt.Printf("%+v\n", timmy)
	fmt.Println(timmy.superpower, (*timmy).superpower)

	// pointers and arrays
	superpowers := &[3]string{"flight", "invisibility", "super strength"}

	fmt.Println(superpowers[0], (*superpowers)[0])
	fmt.Println(superpowers[1:2], (*superpowers)[1:2])

	// pointers as params
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        41,
	}
	fmt.Printf("%+v\n", rebecca)
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)
	rebecca.birthday()
	fmt.Printf("%+v\n", rebecca)

	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry)

	// internal pointers
	player := character{name: "Matt"}
	levelUp(&player.stats)
	fmt.Printf("%+v\n", player)

	// pointers and arrays
	var board [8][8]rune
	fmt.Printf("= %c \n", board[0][0])
	reset(&board)
	fmt.Printf("= %c \n", board[0][0])

	// maps is pointers
	// wrong:
	// func demolish(planets *map[string]string)

	// slices is pointers structure (pointer+cap+len)

	//
	planets := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	reclassify(&planets)
	fmt.Println(planets)

	// pointers and interfaces
	shout(martian{})
	shout(&martian{}) // pointers have the same interface

	pew := laser(2)
	shout(&pew)
	// shout(pew) // wrong

	t := turtle{0, 0}
	t.where()
	t.up()
	t.where()
	t.right()
	t.where()
	t.down()
	t.where()
	t.left()
	t.where()
}
