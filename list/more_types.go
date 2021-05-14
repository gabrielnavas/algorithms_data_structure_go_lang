package main

import (
	"container/list"
	"fmt"
)

func main() {
	var value interface{}
	var intList list.List

	// some types examples
	intList.PushBack(2)
	intList.PushBack(4.1)
	intList.PushBack("Gabriel")
	intList.PushBack(true)

	for element := intList.Front(); element != nil; element = element.Next() {
		value = element.Value
		fmt.Println(value)
	}

	/*
		OUTPUT:
		2
		4.1
		Gabriel
		true
	*/
}
