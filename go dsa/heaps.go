package main

import "fmt"

// MaxHeap struct holds a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert method for adding new elements
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract method to remove and return the largest element
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	last := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Cannot extract from empty array")
		return -1
	}
	h.array[0] = h.array[last]
	h.array = h.array[:last]
	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp that rearranges the max heap after a change
func (h *MaxHeap) maxHeapifyUp(indx int) {
	for h.array[parent(indx)] < h.array[indx] {
		h.swap(parent(indx), indx)
		indx = parent(indx)
	}
}

// maxHeapifyDown that rearranges the max heap after an extraction
func (h *MaxHeap) maxHeapifyDown(indx int) {
	last := len(h.array) - 1
	l, r := left(indx), right(indx)
	child2Compare := 0
	for l <= last {
		if l == last { // left child is only child
			child2Compare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			child2Compare = l
			child2Compare = r
		}
		// compare array value of current index to the larger child and swap if smaller
		if h.array[indx] < h.array[child2Compare] {
			h.swap(indx, child2Compare)
			indx = child2Compare
			l, r = left(indx), right(indx)
		} else {
			return
		}
	}
}

// parent returns the index of the parent
func parent(i int) int {
	return (i - 1) / 2
}

// left return the index of  left child
func left(i int) int {
	return 2*i + 1
}

// returns the index of the right child
func right(i int) int {
	return 2*i + 2
}

// swap interchanges the indices of two elements
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 14, 15, 18}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
