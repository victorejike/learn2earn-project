package main

import "strings"

func removeToken(tokens []string, index int) []string {
	return append(tokens[:index], tokens[index+1:]...)
}

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

func startsWithVowelOrH(word string) bool {
	if len(word) == 0 {
		return false
	}
	first := strings.ToLower(string(word[0]))
	return first == "a" || first == "e" || first == "i" || first == "o" || first == "u" || first == "h"
}
