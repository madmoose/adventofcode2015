package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	floor := 0
	r := bufio.NewReader(os.Stdin)
	for {
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
	}

	fmt.Println(floor)
}
