package main

import "fmt"

type Queue []string

func (s *Queue) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Queue) Enqueue(str string) {
	*s = append(*s, str)
}

func (s *Queue) Dequeue() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		elem := (*s)[0]
		*s = (*s)[1:]
		return elem, true
	}
}

func main() {
	var queue Queue
	queue.Enqueue("Data Structures")
	queue.Enqueue("And")
	queue.Enqueue("Algorithms")
	queue.Enqueue("Tutorial")
	queue.Enqueue("Using Golang")

	for !queue.IsEmpty() {
		x, y := queue.Dequeue()
		if y == true {
			fmt.Println(x)
		}
	}
}
