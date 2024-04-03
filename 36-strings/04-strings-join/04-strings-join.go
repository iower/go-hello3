package main

import (
	"fmt"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE %s"

var refStringSlice = []string{
	"FIRST_NAME = 'Jack'",
	"INSURANCE_NO = 123",
	"EFFECTIVE_FROM = SYSDATE",
}

type JoinFunc func(piece string) string

func JoinWithFunc(refStringSlice []string, joinFunc JoinFunc) string {
	result := refStringSlice[0]
	for _, val := range refStringSlice[1:] {
		result += joinFunc(val) + val
	}
	return result
}

func main() {
	sentence := strings.Join(refStringSlice, " AND ")
	fmt.Printf(selectBase+"\n", sentence)

	jF := func(p string) string {
		if strings.Contains(p, "INSURANCE") {
			return " OR "
		}
		return " AND "
	}

	sentence2 := JoinWithFunc(refStringSlice, jF)
	fmt.Printf(selectBase+"\n", sentence2)
}
