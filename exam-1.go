package main

import "fmt"

func main() {
	a := 6
	a = a + 1
	fmt.Println(a)
}

func add(n int) int {
	return n + 1
}
