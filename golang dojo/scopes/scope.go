package main

import "fmt"

var fiveGlobal int = 5

func main() {
	// SCOPES IN GO
	{
		var integer int = 430
		fmt.Println("inside scope", integer)
	}
	var integer int = 320

	fmt.Println("The integer is", integer)

	fmt.Println("Five in the global scope variable is", fiveGlobal)

	fmt.Println("These 3s are declared in another file, but in the same package", Three)

}
