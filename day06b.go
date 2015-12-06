package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lights [1000][1000]int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var x0, y0, x1, y1 int
		if n, _ := fmt.Sscanf(line, "turn on %d,%d through %d,%d\n", &x0, &y0, &x1, &y1); n == 4 {
			for y := y0; y <= y1; y++ {
				for x := x0; x <= x1; x++ {
					lights[y][x]++
				}
			}
		} else if n, _ := fmt.Sscanf(line, "turn off %d,%d through %d,%d\n", &x0, &y0, &x1, &y1); n == 4 {
			for y := y0; y <= y1; y++ {
				for x := x0; x <= x1; x++ {
					if lights[y][x] > 0 {
						lights[y][x]--
					}
				}
			}
		} else if n, _ := fmt.Sscanf(line, "toggle %d,%d through %d,%d\n", &x0, &y0, &x1, &y1); n == 4 {
			for y := y0; y <= y1; y++ {
				for x := x0; x <= x1; x++ {
					lights[y][x] += 2
				}
			}
		}
	}

	brightness := 0
	for y := 0; y != 1000; y++ {
		for x := 0; x != 1000; x++ {
			brightness += lights[y][x]
		}
	}

	fmt.Println(brightness)
}
