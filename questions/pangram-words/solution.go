package main

import (
	"fmt"
)

func findUniqueStrings(words []string) ([]string, bool) {
	usedWords := []string{}
	charSet := make(map[rune]bool)
	return backtrack(words, usedWords, charSet, 0)
}

func backtrack(words []string, usedWords []string, charSet map[rune]bool, start int) ([]string, bool) {
	if len(usedWords) == 5 {
		return usedWords, true
	}
	for i := start; i < len(words); i++ {
		word := words[i]
		wordSet := make(map[rune]bool)
		unique := true
		for _, ch := range word {
			if charSet[ch] || wordSet[ch] {
				unique = false
				break
			}
			wordSet[ch] = true
		}
		if unique {
			for ch := range wordSet {
				charSet[ch] = true
			}
			usedWords = append(usedWords, word)
			result, found := backtrack(words, usedWords, charSet, i+1)
			if found {
				return result, true
			}
			for ch := range wordSet {
				delete(charSet, ch)
			}
			usedWords = usedWords[:len(usedWords)-1]
		}
	}
	return nil, false
}

func main() {
	words := []string{"abcde", "fghij", "klmnf", "pqrst", "ghijo", "uvwxy"}
	result, found := findUniqueStrings(words)
	if found {
		fmt.Println("Found strings with 25 unique characters:")
		for _, word := range result {
			fmt.Println(word)
		}
	} else {
		fmt.Println("Could not find 5 strings with unique characters")
	}
}
