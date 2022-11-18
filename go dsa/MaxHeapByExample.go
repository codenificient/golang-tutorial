package main

import "fmt"

type maxHeap struct {
	heapArry []int
	size     int
	maxSize  int
}

func newMaxHeap(size int) *maxHeap {
	maxheap := &maxHeap{
		heapArry: []int{},
		size:     0,
		maxSize:  size,
	}
	return maxheap
}

func (m *maxHeap) leaf(indx int) bool {
	if indx >= (m.size/2) && indx <= m.size {
		return true
	}
	return false
}

func (m *maxHeap) parent(indx int) int {
	return (indx - 1) / 2
}

func (m *maxHeap) left(indx int) int {
	return indx*2 + 1
}

func (m *maxHeap) right(indx int) int {
	return indx*2 + 2
}

func (m *maxHeap) swap(itm1, itm2 int) {
	m.heapArry[itm1], m.heapArry[itm2] = m.heapArry[itm2], m.heapArry[itm1]
}

func (m *maxHeap) insert(item int) error {
	if m.size >= m.maxSize {
		return fmt.Errorf("heap is full")
	}
	m.heapArry = append(m.heapArry, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}

func (m *maxHeap) upHeapify(indx int) {
	for m.heapArry[indx] > m.heapArry[m.parent(indx)] {
		m.swap(indx, m.parent(indx))
		indx = m.parent(indx)
	}
}

func (m *maxHeap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	largest := current
	leftChildIndx, rightChildIndx := m.left(current), m.right(current)
	// If current is smallest then return
	if leftChildIndx < m.size && m.heapArry[leftChildIndx] > m.heapArry[largest] {
		largest = leftChildIndx
	}
	if rightChildIndx < m.size && m.heapArry[rightChildIndx] > m.heapArry[largest] {
		largest = rightChildIndx
	}
	if largest != current {
		m.swap(current, largest)
		m.downHeapify(largest)
	}
	return
}

func (m *maxHeap) buildMaxHeap() {
	for index := (m.size / 2) - 1; index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *maxHeap) pop() int {
	top := m.heapArry[0]
	m.heapArry[0] = m.heapArry[m.size-1]
	m.size--
	m.downHeapify(0)
	return top
}

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	testHeap := newMaxHeap(len(inputArray))
	for _, item := range inputArray {
		testHeap.insert(item)
	}
	fmt.Print("The max heap is ")
	for _ = range inputArray {
		fmt.Print(testHeap.pop(), "  ")
	}
}
