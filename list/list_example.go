package main

import (
	"container/list"
	"fmt"
)

func main() {
	var value int
	var intList list.List
	intList.PushBack(2)
	intList.PushBack(4)
	intList.PushBack(5)
	intList.PushBack(10)

	for element := intList.Front(); element != nil; element = element.Next() {
		value = element.Value.(int)
		fmt.Println(value)
	}

	/*
		OUTPUT:
		2
		4
		5
		10
	*/
}
