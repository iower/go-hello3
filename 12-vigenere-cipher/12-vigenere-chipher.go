package main

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := "ECFRZKYGLGRMUSDHRXK"
	keyword := "GOLANG"

	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i]
		k := keyword[i%len(keyword)]

		res := c + 'A' - k
		if res < 'A' {
			res += 26
		} else if res > 'Z' {
			res -= 26
		}

		fmt.Printf("%c %c %c \n", c, k, res)
	}

	fmt.Println("---------")

	plainText := "your message goes here"
	plainText2 := strings.Replace(plainText, " ", "", -1)
	plainText3 := strings.ToUpper(plainText2)

	for i := 0; i < len(plainText3); i++ {
		p := plainText3[i]
		k := keyword[i%len(keyword)]

		res := p + k - 'A'
		if res < 'A' {
			res += 26
		} else if res > 'Z' {
			res -= 26
		}

		fmt.Printf("%c %c %c \n", p, k, res)
	}
}
