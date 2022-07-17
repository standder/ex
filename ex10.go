package main

import "fmt"

type cat struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello world")
	var a int = 10
	fmt.Println(a)
	newCat := cat{name: "egg", age: 10}
	fmt.Println(newCat)
}
