package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	i := 0
	for ; ; i++ {
		b := md5.Sum([]byte("iwrupvqb" + strconv.Itoa(i)))
		if b[0] == 0 && b[1] == 0 && b[2]&0xf0 == 0 {
			break
		}
	}

	fmt.Println(i)
}
