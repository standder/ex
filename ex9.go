package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	a := boy{name: "hugo", age: 21, sex: "male"}
	fmt.Println("I'm a boy")
	fmt.Println("This is my information")
	fmt.Println(a)
	b := 10
	c := 12
	fmt.Println(add(b, c))
}

type boy struct {
	name string
	age  int
	sex  string
}

func add(x int, y int) int {
	return x + y
}
