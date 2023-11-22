package main

import (
	"fmt"
	"os"
)

const (
	rows, cols = 9, 9
)

type Cell struct {
	num   int8
	fixed bool
}

type Grid [][]Cell

func NewSudoku(nums [rows][cols]int8) *Grid {

}

func (g *Grid) Set(row, col int, num int8) error {
	return nil
}
func (g *Grid) Clear(row, col int8) error {
	return nil
}

func main() {
	s := NewSudoku([rows][cols]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(1, 1, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range s {
		fmt.Println(row)
	}
}
