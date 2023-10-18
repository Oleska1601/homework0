package main

import (
	"fmt"
	"strings"
)

// функция для проверки введенной строки на палиндом
func IsPalindrome(input string) bool {
	words := []rune(strings.ReplaceAll(strings.ToLower(input), " ", ""))
	length := len(words)
	for i := 0; i < length/2; i++ {
		if words[i] != words[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("А роза упала на лапу Азора"))
	fmt.Println(IsPalindrome("оаоао"))
	fmt.Println(IsPalindrome("hello world"))
}
