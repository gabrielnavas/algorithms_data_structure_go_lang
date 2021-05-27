package main

import "fmt"

func main() {
	dtoFactory := NewDataTransferObjectFactory()

	other_dto := dtoFactory.getDataTransferObject("other_dto")
	if other_dto != nil {
		fmt.Println(other_dto.getId())
	} else {
		fmt.Printf("dto is nil: %v\n", other_dto)
	}

	customer := dtoFactory.getDataTransferObject("customer")
	if customer != nil {
		fmt.Printf("ID Customer is: %s\n", customer.getId())
	}

	address := dtoFactory.getDataTransferObject("address")
	if address != nil {
		fmt.Printf("ID address is : %s\n", address.getId())
	}
}
