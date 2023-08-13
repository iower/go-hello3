package main

import "fmt"

func main() {
	peace1 := "peace"
	var peace2 = "peace"
	var peace3 string = "peace"
	fmt.Println(peace1, peace2, peace3)

	var blank string
	fmt.Println(blank)

	fmt.Println("esc\nape")
	fmt.Println(`no esc\nape`)

	fmt.Println(`
		peace be upon you
		upon you be peace
	`)

	// the same type

	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)

	// rune, byte
	type byte2 = uint8
	type rune2 = int32

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	grade := 'A'
	var grade2 = 'A'
	var grade3 rune = 'A'
	var star byte = '*'
	fmt.Printf("%v %v %v %v\n", grade, grade2, grade3, star)
	fmt.Printf("%c %c %c %c\n", grade, grade2, grade3, star)

}
