package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasSplitPair(s string) bool {
	for n := 2; n < len(s); n++ {
		if s[n] == s[n-2] {
			return true
		}
	}

	return false
}

func hasRepeatedCouple(s string) bool {
	for n := 3; n < len(s); n++ {
		for m := 1; m < n-1; m++ {
			if s[n] == s[m] && s[n-1] == s[m-1] {
				return true
			}
		}
	}

	return false
}

func isNice(word string) bool {
	return hasSplitPair(word) &&
		hasRepeatedCouple(word)
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
