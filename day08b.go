package main

import (
	"bufio"
	"fmt"
	"os"
)

func escapedLength(s string) int {
	length := 2
	for _, r := range s {
		switch r {
		case '"', '\'', '\\':
			length += 2
		default:
			length += 1
		}
	}

	return length
}

func main() {
	total := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		total += escapedLength(line) - len(line)
	}

	fmt.Println(total)
}
