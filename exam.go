package main

import "fmt"

func main() {
	a := 6
	add(&a)
	fmt.Println(a)
}

func add(n *int) {
	*n = *n + 1
}
