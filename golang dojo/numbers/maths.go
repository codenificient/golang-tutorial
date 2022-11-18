package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Now we try the Math library")
	//A WHOLE BUNCH OF INTEGERS
	var x = 12
	var x32 int32 = 324
	var x64 int64 = 492

	fmt.Printf("x type is %T", x)
	fmt.Println()
	fmt.Printf("x32 type is %T", x32)
	fmt.Println()
	fmt.Printf("x64 type is %T", x64)
	var xuint uint = 3
	fmt.Println()
	fmt.Printf("The following integer is %T", xuint)

	//A WHOLE BUNCH OF FLOATING POINT NUMBERS
	pi := 3.1415
	fmt.Println("The current value of PI is", pi)
	fmt.Printf("And the type of pi is %T", pi)
	fmt.Println()
	var pi32 float32 = 3.1415
	fmt.Printf("The float 32 version of pi has type %T", pi32)
	fmt.Println()
	var pi64 = float64(pi32)
	fmt.Printf("After casting to float 64 version of pi has type %T", pi64)

	//CONVERSION BETWEEN FLOAT AND INTEGERS
	a := 1
	b := 3.1415
	fmt.Println()
	fmt.Println("a is", a)
	fmt.Printf("Type of a is %T", a)
	fmt.Println()
	fmt.Println("b is", b)
	fmt.Printf("Type of a is %T", b)

	fmt.Println()
	a = int(b)
	fmt.Printf("After conversion, a is now of type %T", a)
	fmt.Println()
	fmt.Println()
	//USING SOME BUILTIN FUNCTIONS IN THE MATH LIBRARY
	fmt.Println("NOW COMES THE MATH LIBRARY")
	fmt.Println(math.Ceil(3.000001))           // 4
	fmt.Println(math.Min(3.000001, -3.000001)) // -3.000001
	fmt.Println(math.Max(3.000001, -3.000001)) // 3.000001
	fmt.Println(math.Abs(-3.000001))           // 3.000001
	fmt.Println(math.Sqrt(7))                  // 4
	fmt.Println(math.Pow(3.000001, 2))         // 10
}
