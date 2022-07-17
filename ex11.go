package main

import (
	"fmt"
	"math"
)

func main() {
	pi := math.Pi
	fmt.Printf("%v %g (%6.2f) %e %E\n", pi, pi, pi, pi, pi)

	integer := 23
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	var trith bool = true
	fmt.Printf("%v %t\n", trith, trith)

	point := 101 + 11i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)

	a := 42
	fmt.Printf("%v %d %x %o %b\n", a, a, a, a, a)
}
