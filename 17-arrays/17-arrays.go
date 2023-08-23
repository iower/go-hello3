package main

import (
	"fmt"
	"strings"
)

func main() {
	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(planets)

	fmt.Println(len(planets))
	fmt.Println(planets[3])
	fmt.Println(planets[3] == "")

	var nums [3]int
	fmt.Println("nums", nums)

	// out of range error
	/*
		for i := 0; i <= 3; i++ {
			fmt.Println(nums[i])
		}
	*/

	// composite literal
	strings := [4]string{"aaa", "bbb", "ccc"}
	strings2 := [4]string{
		"aaa",
		"bbb",
		"ccc",
		"ddd",
	}
	strings3 := [...]string{
		"aaa",
		"bbb",
		"ccc",
		"ddd",
	}
	fmt.Println(strings)
	fmt.Println(strings2)
	fmt.Println(strings3)
	fmt.Println(len(strings), len(strings2), len(strings3))

	// iteration
	for i := 0; i < len(strings3); i++ {
		str := strings3[i]
		fmt.Println(i, str)
	}

	for i, str := range strings3 {
		fmt.Println(i, str)
	}

	// copy arrays
	strings3copy := strings3
	strings3[0] = "changed"
	fmt.Println(strings3, strings3copy)

	// pass by value (copy value)
	func(strings [4]string) {
		for i := range strings {
			strings[i] += "Changed only inside "
		}
		fmt.Println(strings)
	}(strings3)
	fmt.Println(strings3)

	// arrays of different lengths are different types
	arr1 := [1]int{1}
	arr2 := [2]int{1, 2}
	// arr2 = arr1 // cannot
	fmt.Println(arr1, arr2)

	// arrays from arrays
	var boardTest [8][8]string
	boardTest[0][0] = "r"
	boardTest[0][7] = "r"

	for column := range boardTest[1] {
		boardTest[1][column] = "p"
	}

	fmt.Println(boardTest)

	var matrix [3][3]int
	fmt.Println(matrix)

	// chess

	var board chess

	// black
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'

	for col := range board[1] {
		board[1][col] = 'p'
		board[6][col] = 'p'
	}

	// white
	board[7][0] = 'R'
	board[7][1] = 'N'
	board[7][2] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'
	board[7][5] = 'B'
	board[7][6] = 'N'
	board[7][7] = 'R'

	printBoard(board)
}

type chess [8][8]rune

func printBoard(board chess) {
	edge := strings.Repeat("=", len(board)*2+3)
	fmt.Println(edge)
	for _, row := range board {
		fmt.Print("| ")
		for _, col := range row {
			if col == 0 { // empty rune
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", col)
			}
		}
		fmt.Println("|")
	}
	fmt.Println(edge)
}
