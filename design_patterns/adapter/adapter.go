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
