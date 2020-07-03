// switch statements in Go

package main

import "fmt"

func main() {
	yourAge := 5
	
	switch yourAge {
	case 16:
		fmt.Println("Go drive")
	case 18:
		fmt.Println("Go vote")
	default:
		fmt.Println("Go Have Fun")
	}
}
