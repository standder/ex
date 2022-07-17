package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	a[0] = 10
	a[1] = 11
	a[2] = 12
	a = append(a, 13)
	fmt.Println(a)

	numbers := make(map[string]int)
	numbers["sudo"] = 10
	numbers["su"] = 12
	numbers["ss"] = 10
	fmt.Println(numbers)

	c := []int{12, 122, 63, 74, 25}
	for k, v := range c {
		fmt.Println(k, v)
	}

	fmt.Println(domath(10, 12))
}

func domath(x int, y int) (int, int) {
	return x + y, x * y
}
