package main

import (
	"fmt"
)

func nextASCII(b byte) byte {
	if b == 255 {
		return b
	}
	return b + 1
}

func prevASCII(b byte) byte {
	if b != 0 {
		return b - 1
	}
	return b
}

func main() {
	next := nextASCII(byte('b'))
	fmt.Println(string(next))
}
