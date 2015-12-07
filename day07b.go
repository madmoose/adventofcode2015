package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Mnemonic uint8

const (
	MOV Mnemonic = iota
	AND
	OR
	RSHIFT
	LSHIFT
	NOT
)

type Wire struct {
	mnemonic   Mnemonic
	src1, src2 string
}

var (
	wires          = make(map[string]Wire)
	memoizedValues = make(map[string]uint16)
)

func value(dst string) uint16 {
	if v, err := strconv.ParseUint(dst, 10, 16); err == nil {
		return uint16(v)
	}
	if v, ok := memoizedValues[dst]; ok {
		return v
	}

	var w Wire = wires[dst]
	var v uint16

	switch w.mnemonic {
	case MOV:
		v = value(w.src1)
	case AND:
		v = value(w.src1) & value(w.src2)
	case OR:
		v = value(w.src1) | value(w.src2)
	case LSHIFT:
		v = value(w.src1) << value(w.src2)
	case RSHIFT:
		v = value(w.src1) >> value(w.src2)
	case NOT:
		v = ^value(w.src1)
	}

	memoizedValues[dst] = v
	return v
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var (
			mnemonic        Mnemonic
			src1, src2, dst string
		)

		if n, _ := fmt.Sscanf(line, "%s -> %s\n", &src1, &dst); n == 2 {
			mnemonic = MOV
		} else if n, _ := fmt.Sscanf(line, "%s AND %s -> %s\n", &src1, &src2, &dst); n == 3 {
			mnemonic = AND
		} else if n, _ := fmt.Sscanf(line, "%s OR %s -> %s\n", &src1, &src2, &dst); n == 3 {
			mnemonic = OR
		} else if n, _ := fmt.Sscanf(line, "%s RSHIFT %s -> %s\n", &src1, &src2, &dst); n == 3 {
			mnemonic = RSHIFT
		} else if n, _ := fmt.Sscanf(line, "%s LSHIFT %s -> %s\n", &src1, &src2, &dst); n == 3 {
			mnemonic = LSHIFT
		} else if n, _ := fmt.Sscanf(line, "NOT %s -> %s\n", &src1, &dst); n == 2 {
			mnemonic = NOT
		} else {
			panic(line)
		}

		wires[dst] = Wire{mnemonic, src1, src2}
	}

	a := value("a")

	memoizedValues = make(map[string]uint16)
	memoizedValues["b"] = a

	fmt.Println(value("a"))
}
