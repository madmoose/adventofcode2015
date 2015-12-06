package main

import "fmt"

func main() {
	ribbon := 0
	for {
		var l, w, h int
		if n, _ := fmt.Scanf("%dx%dx%d\n", &l, &w, &h); n != 3 {
			break
		}

		if l < h && w < h {
			ribbon += l + l + w + w
		} else if l < w {
			ribbon += l + l + h + h
		} else {
			ribbon += w + w + h + h
		}

		ribbon += l * w * h

	}

	fmt.Println(ribbon)
}
