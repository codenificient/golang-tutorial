package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	fmt.Println(mp)
	mp["banana"] = 90

	fmt.Println(mp)

	tim, ok := mp["tim"]
	fmt.Println(tim, ok)
}
