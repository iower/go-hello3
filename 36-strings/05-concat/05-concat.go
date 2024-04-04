package main

import (
	"bytes"
	"fmt"
)

func main() {
	strings := []string{"This", "is", "even", "more", "perfomant"}

	// concat buffer
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}
	fmt.Println(buffer.String())

	// concat copy
	bs := make([]byte, 100)
	bl := 0

	for _, val := range strings {
		bl += copy(bs[bl:], []byte(val))
	}
	fmt.Println(string(bs[:]))
}
