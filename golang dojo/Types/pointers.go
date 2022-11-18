package main

import "fmt"

func main() {
	fmt.Println("We will see pointers now")

	// original way
	var x1 int = 1
	var p1 *int = &x1

	// new way
	x := 1
	p := &x

	fmt.Println("The value of the integer is", x1)
	fmt.Println("The value of the integer pointer is", p1)

	fmt.Println("The value of the integer is", x)
	fmt.Println("The value of the integer pointer is", p)

	// ** TYPE DECLARATIOSN IN GO **

	// 	fahrenheit and celcius
	type fahrenheit int
	type celcius int 

	var f fahrenheit = 103
	var c celcius = 0

	fmt.Println(f, c)

	c = celcius((f - 32) * 5 / 9)

	fmt.Println(c)
}