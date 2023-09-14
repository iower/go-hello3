package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

var percent = 0.0

func (u Universe) Seed() {
	allN := 0.0
	liveN := 0.0
	for row := range u {
		for col := range u[row] {
			if rand.Intn(4) == 0 {
				u[row][col] = true
				liveN++
			} else {
				u[row][col] = false
			}
			allN++
		}
	}
	percent = math.Round(liveN / allN * 100)
}

func (u Universe) isAlive(x, y int) bool {
	xMod := (x + width) % width
	yMod := (y + height) % height
	return u[yMod][xMod]
}

func (u Universe) Neighbors(x, y int) int {
	bools := []bool{
		u.isAlive(x-1, y-1), u.isAlive(x, y-1), u.isAlive(x+1, y-1),
		u.isAlive(x-1, y) /* the same cell */, u.isAlive(x+1, y),
		u.isAlive(x-1, y+1), u.isAlive(x, y+1), u.isAlive(x+1, y+1),
	}
	n := 0
	for _, val := range bools {
		if val {
			n++
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.isAlive(x, y)
}

func (u Universe) Set(x, y int, val bool) {
	u[y][x] = val
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func PrintBorder(u Universe) {
	fmt.Println()
}

func (u Universe) Show() {
	hr := strings.Repeat("#", len(u[0])+2) + "\n"
	buf := ""
	buf += hr
	for _, row := range u {
		buf += "#"
		for _, col := range row {
			if col {
				buf += "O"
			} else {
				buf += " "
			}
		}
		buf += "#\n"
	}
	buf += hr
	buf += fmt.Sprintf("Game of Life | Seeded %.0f%% \n", percent)
	fmt.Print("\033[H\033[2J", buf)
}

func main() {
	u1, u2 := NewUniverse(), NewUniverse()
	u1.Seed()

	for true {
		u1.Show()
		Step(u1, u2)
		u1, u2 = u2, u1
		time.Sleep(time.Second / 15)
	}
}
