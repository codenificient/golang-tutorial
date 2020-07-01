package main

import "fmt"

func main() {
	ourNewFunction()
}

func ourNewFunction() {
	fmt.Println("This is a function being called by main!")
	anotherFunction()
}

func anotherFunction() {
	fmt.Println("This is another function being called by ourNewFunction! which will be later called by main!")
}
