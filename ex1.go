package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

func init() {
	_, num, str := number()
	fmt.Println(num, str)
}

func number() (int, int, string) {
	a, b, c := 1, 2, "su"
	return a, b, c
}
