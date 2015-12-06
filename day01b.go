package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	floor := 0
	pos := 1
	r := bufio.NewReader(os.Stdin)

	for ; ; pos++ {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}

		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor < 0 {
			break
		}
	}

	fmt.Println(pos)
}
