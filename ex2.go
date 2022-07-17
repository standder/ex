package main

import "fmt"

func main() {
	const (
		WEIGHT int = 20
		HIGHT  int = 10
		LONGER int = 30
	)
	var area int = WEIGHT * HIGHT * LONGER
	fmt.Println(area)
}
