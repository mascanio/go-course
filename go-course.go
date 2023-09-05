package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "τ = 2π"
	var reverse []byte

	for i := len(input); i > 0; {
		r, s := utf8.DecodeLastRuneInString(input[:i])
		i -= s
		reverse = utf8.AppendRune(reverse, r)
	}
	fmt.Println(string(reverse))
}
