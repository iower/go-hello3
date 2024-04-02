package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {
	lookFor := "lamb"
	contain := strings.Contains(refString, lookFor)
	fmt.Println(contain)

	lookFor = "wolf"
	contain = strings.Contains(refString, lookFor)
	fmt.Println(contain)

	lookFor = "Mary"
	contain = strings.HasPrefix(refString, lookFor)
	fmt.Println(contain)

	lookFor = "lamb"
	contain = strings.HasSuffix(refString, lookFor)
	fmt.Println(contain)

	lookFor = "Mary"
	i := strings.Index(refString, lookFor)
	fmt.Println(i)

	lookFor = "!"
	i = strings.Index(refString, lookFor)
	fmt.Println(i)
}
