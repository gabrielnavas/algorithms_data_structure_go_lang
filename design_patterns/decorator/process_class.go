package main

import "fmt"

type ProcessClass struct{}

func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}
