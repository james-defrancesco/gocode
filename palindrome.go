package main

import (
	"fmt"
	"strings"
)

func main() {

	var ip string
	fmt.Println("Enter string: ")
	fmt.Scanf("%s\n", &ip)
	ip = strings.ToLower(ip)
	fmt.Println(isP(ip))
}

//Function to test if the string entered is a Palindrome

func isP(s string) string {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return "NO. It's not a Palindrome."
		}
	}
	return "YES! You've entered a Palindrome"
}
