package main

import "fmt"

type Customer struct {
	name string
	id   int
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("Create a customer")
	customer.name = name
	return customer
}
