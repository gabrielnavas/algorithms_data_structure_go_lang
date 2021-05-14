package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *IntegerHeap) Push(elementint interface{}) {
	*iheap = append(*iheap, elementint.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	var previous IntegerHeap = *iheap
	len := len(previous)
	element := previous[len-1]
	*iheap = previous[0 : len-1]
	return element
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{4, 2, 5}
	heap.Init(intHeap) // need init for
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Println(heap.Pop(intHeap))
	}
}
