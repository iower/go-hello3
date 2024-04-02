package main

import (
	"fmt"
	"regexp"
	"strings"
)

const refString = "Mary had a little lamb"

func printWords(words []string) {
	fmt.Println("---")
	for id, word := range words {
		fmt.Printf("Word %d is: %s\n", id, word)
	}
	fmt.Println("---")
}

func main() {
	words := strings.Fields(refString)
	printWords(words)

	words = strings.Split(refString, "a")
	printWords(words)

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("_-", r)
	}

	const refString = "Mary_had-a_little-lamb"

	words = strings.FieldsFunc(refString, splitFunc)
	printWords(words)

	words = regexp.MustCompile("[_-]{1}").Split(refString, -1)
	printWords(words)
}
