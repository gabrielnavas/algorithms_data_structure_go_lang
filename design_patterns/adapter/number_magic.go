package main

import "fmt"

//Adaptee Struct
type NumberMagic struct {
	number int
}

// Adaptee class method convert
func (numMagic NumberMagic) convert() int {
	numMagic.number = 1
	fmt.Println("number magic maked")
	return numMagic.number
}
