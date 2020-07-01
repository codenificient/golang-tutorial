package main

import "fmt"

const ourString string = "This is a function being called by main!"

func main() {
	ourNewFunction()
}

func ourNewFunction() {
	fmt.Println(ourString)
	anotherFunction()
}

func anotherFunction() {
	fmt.Println(ourString + " but accessible from a new function")
	fmt.Println("This is another function being called by ourNewFunction! which will be later called by main!")
}
