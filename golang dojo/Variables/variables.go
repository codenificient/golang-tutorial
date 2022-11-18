package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Variables")
	
	var integer int = 32

	var integer2, integer3 int = 3, 42

	fmt.Println("The integer value is", integer)
	fmt.Println("The integer values are", integer2, "and", integer3)

	fmt.Println("We will see booleans now")

	// setting variables with := expression

	boolean := false

	fmt.Println("We will see booleans now")
	fmt.Println("Currently the boolean value is", boolean)
}