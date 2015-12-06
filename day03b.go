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
	var (
		houses  = make(map[sleigh]int)
		sleighs = [2]sleigh{}
		i       int
		err     error
	)

	for i, _ = range sleighs {
		houses[sleighs[i]]++
	}

	i = 0
	r := bufio.NewReader(os.Stdin)
	for {
		var c rune
		if c, _, err = r.ReadRune(); err != nil {
			break
		}

		sleighs[i].move(c)
		houses[sleighs[i]]++
		i = (i + 1) % len(sleighs)
	}

	fmt.Println(len(houses))
}
