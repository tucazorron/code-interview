package main

import (
	"fmt"
)

func solution(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	charCount := make(map[rune]int)
	for _, char := range word1 {
		charCount[char]++
	}
	for _, char := range word2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solution("abc", "cba")) // true
	fmt.Println(solution("abc", "cb"))  // false
	fmt.Println(solution("abc", "cbb")) // false
}
