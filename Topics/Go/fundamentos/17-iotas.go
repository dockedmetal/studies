package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
)

func main() {
	fmt.Println(a, b)

	c := 1

	fmt.Println(b == c) // true
}
