package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Node struct {
	From, To string
}

var distance = make(map[Node]int)

func pathLength(path []string) int {
	length := 0
	for i := 0; i < len(path)-1; i++ {
		length += distance[Node{path[i], path[i+1]}]
	}
	return length
}

func reverse(p sort.Interface, begin, end int) {
	end--
	for begin < end {
		p.Swap(begin, end)
		begin++
		end--
	}
}

func nextPermutation(p sort.Interface) bool {
	if p.Len() <= 1 {
		return false
	}

	i := p.Len() - 1
	for {
		i2 := i
		i--
		if p.Less(i, i2) {
			j := p.Len() - 1
			for !p.Less(i, j) {
				j--
			}
			p.Swap(i, j)
			reverse(p, i2, p.Len())
			return true
		}
		if i == 0 {
			reverse(p, 0, p.Len())
			return false
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	names := make(map[string]struct{})

	for scanner.Scan() {
		line := scanner.Text()

		var (
			from, to string
			dist     int
		)

		if n, _ := fmt.Sscanf(line, "%s to %s = %d\n", &from, &to, &dist); n != 3 {
			log.Panic(line)
		}

		distance[Node{from, to}] = dist
		distance[Node{to, from}] = dist
		names[from] = struct{}{}
		names[to] = struct{}{}
	}

	path := make([]string, 0, len(names))
	for name := range names {
		path = append(path, name)
	}
	sort.Sort(sort.StringSlice(path))

	shortest := math.MaxInt32
	for nextPermutation(sort.StringSlice(path)) {
		length := pathLength(path)
		if length < shortest {
			shortest = length
		}
	}

	fmt.Println(shortest)
}
