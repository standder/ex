package main

const (
	a = iota
	b
	c
	d = "ha"
	e
	f
	g = 101
	h = iota
	i
)

func main() {
	println(a, b, c, d, e, f, g, h, i)
}
