package main

import "fmt"

//IProcess interface
type IProcess interface {
	process() string
}

//Adapter struct
type AdapterIntToString struct {
	numberMagic NumberMagic
}

//Adapter class method process
func (adapt AdapterIntToString) process() string {
	fmt.Println("Adapter Int to String process")
	num := adapt.numberMagic.convert()
	return fmt.Sprintf("%d", num)
}

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
