package main

import (
	"fmt"
)

func main() {
	const (
		a = 1 << iota
		b = 2 << iota
		c
		d
		e
	)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
}
