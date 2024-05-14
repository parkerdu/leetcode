package main

import (
	"fmt"
)

func reverseWords(s string) string {
	var res string
	sigle := make([]byte, len(s))
	left, right := len(sigle)-1, len(sigle)-1
	for i := range s {
		if s[i] != ' ' {
			sigle[left] = s[i]
			left--
		} else {
			res += string(sigle[left+1:right+1]) + " "
			right = left
		}
	}
	return res + string(sigle[left+1:right+1])
}

func main() {
	s := "let's go"
	fmt.Println(reverseWords(s))
}
