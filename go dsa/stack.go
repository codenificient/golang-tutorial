package main

import "fmt"


type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		idx := len(*s) - 1
		elem := (*s)[idx]
		*s = (*s)[:idx]
		return elem, true
	}
}

func main() {
	var stack Stack
	stack.Push("Using Golang")
	stack.Push("Tutorial")
	stack.Push("Algorithms")
	stack.Push("And")
	stack.Push("Data Structures")


	for !stack.IsEmpty() {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}