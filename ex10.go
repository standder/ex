package main

import "fmt"

func main() {
	fmt.Println(add(10, 12))
	a := cat{name: "kitty", age: 12, sex: "male"}
	fmt.Println(a)
}

func add(x int, y int) (int, int) {
	return x + y, x * y
}

type cat struct {
	name string
	age  int
	sex  string
}
