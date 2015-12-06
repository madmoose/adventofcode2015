package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hasThreeWovels(s string) bool {
	n := 0
	for _, r := range s {
		if strings.ContainsRune("aeiou", r) {
			n++
		}
	}

	return n >= 3
}

func hasDoubleLetter(s string) bool {
	for n := 1; n < len(s); n++ {
		if s[n] == s[n-1] {
			return true
		}
	}

	return false
}

func hasNoBadCombinations(s string) bool {
	bads := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}

	for _, b := range bads {
		if strings.Contains(s, b) {
			return false
		}
	}

	return true
}

func isNice(word string) bool {
	return hasThreeWovels(word) &&
		hasDoubleLetter(word) &&
		hasNoBadCombinations(word)
}

func main() {
	nice := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		if isNice(word) {
			nice++
		}
	}

	fmt.Println(nice)
}
