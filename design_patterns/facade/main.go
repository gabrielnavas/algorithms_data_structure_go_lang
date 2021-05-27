package main

import "fmt"

// the client is main function
func main() {
	facade := NewBranchManagerFacade()
	var customer *Customer
	var account *Account

	// create customer and account
	customer, account = facade.CreateCustomerAccount("Gabriel Smith", "Savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)

	// create a transaction
	transaction := facade.CreateTransaction("21223", "31235", 1000)
	fmt.Println(transaction.amount)
}
