package main

import (
	"fmt"
	"unicode/utf8"
)

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

	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)

	//
	message := "shalom"
	message = "salām"
	fmt.Println(message)
	c := message[5]
	fmt.Printf("%c\n", c)

	// strings are immutable
	// message[5] = 't'

	fmt.Println(len(message))
	for i := 0; i < len(message); i++ {
		fmt.Printf("%c\n", message[i])
	}
	fmt.Println("---")

	// Caesar chipher
	char := 'z'
	fmt.Printf("%c -> ", char)
	char += 3
	if char > 'z' {
		char -= 26
	}
	fmt.Printf("%c \n", char)

	//
	char2 := 'g'
	char2 = char2 - 'a' + 'A'
	fmt.Printf("%c \n", char2)

	// ROT13
	msg := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(msg); i++ {
		c := msg[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	// utf-8
	question := "¿Cómo estás?"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	firstChar, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c - %v bytes \n", firstChar, size)

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}

	// decode Caesar chipher
	encoded := "L fdph, L vdz, L frqtxhuhg"
	for i := 0; i < len(encoded); i++ {
		char := encoded[i]
		if char >= 'a' && char <= 'z' {
			char -= 3
			if char < 'a' {
				char += 26
			}
		} else if char >= 'A' && char <= 'Z' {
			char -= 3
			if char < 'A' {
				char += 26
			}
		}

		fmt.Printf("%c", char)
	}
	fmt.Println()

	//
	mes := "Hola Estación Espacial Internacional"

	for _, c := range mes {
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
