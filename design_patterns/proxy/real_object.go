package main

import "fmt"

type RealObject struct{}

func (ro RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}
