package main

import "fmt"

func main() {
	var ourFirstString string = "This is a literal sentence which is the same thing as our first string"
	fmt.Println(ourFirstString)

	var thisSecondString string = " and we just added some other jank"

	fmt.Println(ourFirstString + thisSecondString)

	var ThirdString string = ourFirstString + " but now added to Our third string"

	fmt.Println(ThirdString)
}
