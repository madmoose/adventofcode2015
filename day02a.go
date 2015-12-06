package main

import "fmt"

func main() {
	area := 0
	for {
		var l, w, h int
		if n, _ := fmt.Scanf("%dx%dx%d\n", &l, &w, &h); n != 3 {
			break
		}

		area += 2*l*w + 2*w*h + 2*l*h

		if l*w < w*h && l*w < l*h {
			area += l * w
		} else if w*h < l*h {
			area += w * h
		} else {
			area += l * h
		}

	}

	fmt.Println(area)
}
