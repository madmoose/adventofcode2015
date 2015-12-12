package main

import "fmt"

func lookAndSay(s string) string {
	var b = make([]byte, 0, 2*len(s))
	var i, j int

	for i = 0; i < len(s); i = j {
		r := s[i]
		for j = i + 1; j < len(s) && s[j] == r; j++ {
		}
		if j-i > 9 {
			panic(j - i)
		}
		b = append(b, byte(48+(j-i)))
		b = append(b, r)
	}

	return string(b)
}

func main() {
	var s string
	if n, err := fmt.Scanf("%s\n", &s); n != 1 {
		panic(err)
	}

	for i := 0; i != 40; i++ {
		s = lookAndSay(s)
	}
	fmt.Println(len(s))
}
