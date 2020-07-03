// if statements in Go

package main

import "fmt"

func main() {
	yourAge := 19

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can enjoy life")
	}
}
