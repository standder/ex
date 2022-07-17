package main

import "fmt"

func main() {
	kitty := cat{name: "kitty", age: 3, sex: "male"}
	kitty.show()
	var a MyInt = 10
	a.add()
	a.show()

}

type cat struct {
	name string
	age  int
	sex  string
}

type MyInt int

func (n *MyInt) add() {
	*n = *n + 1
}

func (n MyInt) show() {
	fmt.Println(n)
}

func (n cat) show() {
	fmt.Println(n)
}
