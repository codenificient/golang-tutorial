package main

import "fmt"

func main() {
	fmt.Println("Now we will see Constants in Go")

	const five int = 5
	fmt.Println("The value of the constant five is", five)

	const pi = 3.1415
	fmt.Println("Constant pi has value of", pi)

	const (
		a = 1
		b = 2
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	const (
		e = 5
		f
		g = 8
		h
	)

	fmt.Println(e, f, g, h)
	fmt.Println("The values of f and g above took the value of their preceding covariables")
}
