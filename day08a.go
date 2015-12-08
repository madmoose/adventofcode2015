package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func unescapedLength(s string) int {
	s = s[1 : len(s)-1]

	length := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			switch s[i+1] {
			case '"', '\'', '\\':
				i += 1
			case 'x':
				i += 3
			default:
				log.Panicf("invalid escape sequence \\%c", s[i+1])
			}
		}
		length++
	}

	return length
}

func main() {
	total := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		total += len(line) - unescapedLength(line)
	}

	fmt.Println(total)
}
