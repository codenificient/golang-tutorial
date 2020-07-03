package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println("for loop; i is now: ", i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println("another for loop; j is now: ", j)
	}
}
