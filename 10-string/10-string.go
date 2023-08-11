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
}
