package main

import "fmt"

type IProcess interface {
	process()
}

// ProcessClass Implementation
type ProcessClass struct{}

func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// Class decorator
type ProcessDecorator struct {
	processInstance *ProcessClass
}

func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		decorator.processInstance.process()
	}
}

func main() {
	process := &ProcessClass{}

	decorator := &ProcessDecorator{}
	// decorator dont have process
	decorator.process()

	// now have
	decorator.processInstance = process
	decorator.process()
}
