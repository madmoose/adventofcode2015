package main

import (
	"bufio"
	"fmt"
	"os"
)

type sleigh struct {
	x int
	y int
}

func (s *sleigh) move(c rune) {
	switch c {
	case '^':
		s.y--
	case '>':
		s.x++
	case 'v':
		s.y++
	case '<':
		s.x--
	}
}

func main() {
	var santa sleigh
	var houses = make(map[sleigh]int)

	houses[santa]++

	r := bufio.NewReader(os.Stdin)
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}

		santa.move(c)
		houses[santa]++
	}

	fmt.Println(len(houses))
}
